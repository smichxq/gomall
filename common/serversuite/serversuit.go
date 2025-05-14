/**
定义一组可复用的 Kitex server.Option
简化各个 RPC 服务的启动配置
配置复用：将所有服务通用的启动参数（元数据、Tracing、Metrics）封装到一个 Suite，各个服务只要引入并调用 Options() 即可。

统一行为：保证所有 RPC 服务在指标采集、Tracing、基础信息上都采用一致的策略，便于监控和调优。
*/

package serversuite

import (
	"github.com/cloudwego/gomall/common/mtl"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"

	// Kitex 官方提供的 Prometheus 拦截器插件
	prometheus "github.com/kitex-contrib/monitor-prometheus"
)

type CommonServerSuite struct {
	CurrentServiceName string
}

func (s CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		// 1. 在 RPC 层加入元数据处理（HTTP2）
		server.WithMetaHandler(transmeta.ClientHTTP2Handler),
		// 2. 设置服务的基本信息（ServiceName，供上报和日志使用）
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		// 3. 挂载 Prometheus Server 端拦截器
		server.WithTracer(prometheus.NewServerTracer(
			"",
			"",
			// 关闭 Kitex 自带的服务端默认监控统计逻辑
			prometheus.WithDisableServer(true),
			prometheus.WithRegistry(mtl.Registry),
		),
		),
	}

	return opts
}
