package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/cloudwego/gomall/app/user/biz/dal"
	"github.com/cloudwego/gomall/app/user/conf"
	"github.com/cloudwego/gomall/common/mtl"
	"github.com/cloudwego/gomall/common/serversuite"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName      = conf.GetConf().Kitex.Service
	RegisterAddr     = conf.GetConf().Registry.RegistryAddress[0]
	ConsulHealthAddr = conf.GetConf().Kitex.ConsulHealthAddr
)

func main() {
	opts := kitexInit()
	// 为了适配服务下线后结束上传指标
	p := mtl.InitTracing(ServiceName)

	// 服务关闭前上传剩余链路数据
	// opentelemetry链路数据分批上传
	defer p.Shutdown(context.Background())

	svr := userservice.NewServer(new(UserServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// 初始化mtl
	// dal与rpc依赖mtl
	mtl.InitMetric(ServiceName, conf.GetConf().Kitex.MetricsPort, RegisterAddr)

	// 加载外部配置
	dal.Init()

	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// 健康检查
	go StartHealthCheckServer(":" + strings.Split(ConsulHealthAddr, ":")[1])

	// server suit
	commonServerSuite := serversuite.CommonServerSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegisterAddr,
		ConsulHealthAddr:   ConsulHealthAddr,
	}
	opts = append(opts, server.WithSuite(commonServerSuite))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}

// 健康监测接口
func StartHealthCheckServer(addr string) {
	go func() {
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		})
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatalf("health check server error: %v", err)
		}
	}()
}
