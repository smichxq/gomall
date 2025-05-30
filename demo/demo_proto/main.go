package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/conf"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi/echoservice"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/middleware"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	consulapi "github.com/hashicorp/consul/api"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	consul "github.com/kitex-contrib/registry-consul"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	opts := kitexInit()

	// 健康检查
	go StartHealthCheckServer(":8899")

	svr := echoservice.NewServer(new(EchoServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// 加载外部配置
	// dal.Init()

	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr), server.WithMiddleware(middleware.Middleware))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	// build a consul register with the consul client
	// 读取配置文件中的注册中心地址(单节点)
	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0], consul.WithCheck(&consulapi.AgentServiceCheck{
		HTTP:                           "http://192.168.3.6:8899/health",
		Interval:                       "1s",
		Timeout:                        "1s",
		DeregisterCriticalServiceAfter: "1m",
	}))
	// r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		log.Fatal("NewConsulRegister", err)
		return
	}

	// 组件注册到服务
	opts = append(opts, server.WithRegistry(r))

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
