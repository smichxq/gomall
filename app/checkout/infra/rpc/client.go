package rpc

import (
	"sync"

	"github.com/cloudwego/gomall/app/checkout/conf"
	checkoutUtils "github.com/cloudwego/gomall/app/checkout/utils"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	once          sync.Once
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
)

func ClientInit() {
	once.Do(func() {
		cartClientInit()
		productClientInit()
		paymentClientInit()
	})
}

func paymentClientInit() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoutUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	checkoutUtils.MustHandleError(err)
}

func productClientInit() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoutUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	checkoutUtils.MustHandleError(err)
}

func cartClientInit() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoutUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	opts = append(opts,
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Kitex.Service}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	CartClient, err = cartservice.NewClient("cart", opts...)
	checkoutUtils.MustHandleError(err)
}
