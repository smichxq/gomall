package clientsuite

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"

	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonClientSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

func (s CommonClientSuite) Options() []client.Option {
	opts := []client.Option{
		// 向注册中心声明调用方服务名称
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		// 使用 HTTP/2 的元数据处理器，允许携带 header 信息等。
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		// 指定调用方使用 gRPC 作为传输协议
		client.WithTransportProtocol(transport.GRPC),
		// kitex-opentelemetry中间件
		client.WithSuite(tracing.NewClientSuite()),
	}

	// 注册中心
	r, err := consul.NewConsulResolver(s.RegistryAddr)
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))

	return opts
}
