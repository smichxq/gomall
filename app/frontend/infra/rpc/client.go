package rpc

import (
	"sync"

	"github.com/cloudwego/gomall/app/frontend/conf"
	frontendUtils "github.com/cloudwego/gomall/app/frontend/utils"
	"github.com/cloudwego/gomall/common/clientsuite"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
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
