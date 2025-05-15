/**
统一管理指标：集中在一个地方创建、维护和暴露所有指标，方便未来扩展自定义业务指标。

与 Consul 集成：让监控端点本身也成为可发现、可健康检查的服务，提高运维水平。

无侵入性：只需在应用入口调用一次 InitMetric，即可全局生效。
*/

package mtl

import (
	"fmt"
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registry *prometheus.Registry

func InitMetric(serviceName, metricsPort, registryAddr string) (registry.Registry, *registry.Info) {
	// 默认的 prometheus.DefaultRegisterer 会和全局指标冲突
	// 这里使用独立的 Registry 避免干扰
	Registry = prometheus.NewRegistry()
	// 注册 Go 运行时指标（goroutine 数、内存分配等）
	Registry.MustRegister(collectors.NewGoCollector())
	// 注册进程级别指标（CPU、文件描述符等）
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	// consul注册
	r, _ := consul.NewConsulRegister(registryAddr)
	// metricsPort转为http地址
	addr, _ := net.ResolveTCPAddr("tcp", "192.168.3.6"+metricsPort)

	fmt.Println("addr: ", addr)
	// 构造注册信息
	registryInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": serviceName},
		// SkipListenAddr: true,
	}

	// 通过 Consul 插件将 metrics 服务(即api为/metrics)也注册到 Consul
	_ = r.Register(registryInfo)

	// 关闭应用程序时取消注册
	server.RegisterShutdownHook(func() {
		r.Deregister(registryInfo)
	})

	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))

	fmt.Println("metricsPort: ", metricsPort)
	// 	启动metrics服务以供Prometheus拉取指标
	go http.ListenAndServe(metricsPort, nil)

	return r, registryInfo
}
