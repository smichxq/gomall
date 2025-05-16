package rpc

import (
	"context"
	"sync"

	"github.com/cloudwego/gomall/app/frontend/conf"
	frontendUtils "github.com/cloudwego/gomall/app/frontend/utils"
	"github.com/cloudwego/gomall/common/clientsuite"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	consulclient "github.com/kitex-contrib/config-consul/client"
	"github.com/kitex-contrib/config-consul/consul"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CheckoutClient checkoutservice.Client
	CartClient     cartservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
	ServiceName    = conf.GetConf().Hertz.Service
	RegisterAddr   = conf.GetConf().Registry.RegistryAddress[0]
	err            error
	opts           = []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegisterAddr,
		}),
	}
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductcatalogserviceClient()
		initCartserviceClient()
		initCheckoutClient()
		initOrderClient()
	})
}

func initUserClient() {
	// 客户端从 Consul 获取服务实例列表已在clientsuite配置
	// 使用对应的IDL客户端
	UserClient, err = userservice.NewClient("user", opts...)
	frontendUtils.MustHandleErr(err)
}

func initProductcatalogserviceClient() {
	// 服务熔断
	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		// "fromServiceName/ToServiceName/method"
		return circuitbreak.RPCInfo2Key(ri)
	})

	// 规则
	cbs.UpdateServiceCBConfig("frontend/product/GetProduct",
		circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2},
	)

	// 熔断器注册
	opts = append(opts, client.WithCircuitBreaker(cbs))

	// 熔断降级注册
	opts = append(opts, client.WithFallback(
		fallback.NewFallbackPolicy(
			fallback.UnwrapHelper(func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
				// 如果没有发生熔断
				if err == nil {
					return resp, nil
				}

				/* 发生熔断 */

				// 根据当前rpc上下文获取熔断的rpc方法
				methodName := rpcinfo.GetRPCInfo(ctx).To().Method()

				// 如果方法不为当前测试的rpc方法
				// 则直接放行
				if methodName != "ListProducts" {
					return resp, err
				}

				// 降级措施
				return &product.ListProductsResp{
					Products: []*product.Product{
						{
							Price:       "6.3",
							Id:          3,
							Picture:     "/static/image/t-shirt.jepg",
							Name:        "1233",
							Description: "test",
						},
						{
							Price:       "6.9",
							Id:          3,
							Picture:     "/static/image/t-shirt.jepg",
							Name:        "1233",
							Description: "test",
						},
						{
							Price:       "6.3",
							Id:          3,
							Picture:     "/static/image/t-shirt.jepg",
							Name:        "1233",
							Description: "test",
						},
					},
				}, nil
			}),
		),
	))

	// 适用于client的配置中心
	consulClient, err := consul.NewClient(consul.Options{
		Addr: RegisterAddr,
	})

	// 注册到suite
	opts = append(opts, client.WithSuite(consulclient.NewSuite("product", ServiceName, consulClient)))

	// 客户端从 Consul 获取服务实例列表已在clientsuite配置
	// 使用对应的IDL客户端
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendUtils.MustHandleErr(err)
}

func initCartserviceClient() {
	// 客户端从 Consul 获取服务实例列表已在clientsuite配置
	// 使用对应的IDL客户端
	CartClient, err = cartservice.NewClient("cart", opts...)
	frontendUtils.MustHandleErr(err)
}

func initCheckoutClient() {
	// 客户端从 Consul 获取服务实例列表已在clientsuite配置
	CheckoutClient, err = checkoutservice.NewClient("checkout", opts...)
	frontendUtils.MustHandleErr(err)
}

func initOrderClient() {
	// 客户端从 Consul 获取服务实例列表已在clientsuite配置
	OrderClient, err = orderservice.NewClient("order", opts...)
	frontendUtils.MustHandleErr(err)
}
