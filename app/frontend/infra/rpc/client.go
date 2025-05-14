package rpc

import (
	"sync"

	"github.com/cloudwego/gomall/app/frontend/conf"
	frontendUtils "github.com/cloudwego/gomall/app/frontend/utils"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient     userservice.Client
	ProductClient  productcatalogservice.Client
	CheckoutClient checkoutservice.Client
	CartClient     cartservice.Client
	once           sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductcatalogserviceClient()
		initCartserviceClient()
		initCheckoutClient()
	})
}

func initUserClient() {
	// 客户端从 Consul 获取服务实例列表
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	frontendUtils.MustHandleErr(err)

	// 使用对应的IDL客户端
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleErr(err)
}

func initProductcatalogserviceClient() {
	// 客户端从 Consul 获取服务实例列表
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	frontendUtils.MustHandleErr(err)

	// 使用对应的IDL客户端
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	frontendUtils.MustHandleErr(err)
}

func initCartserviceClient() {
	// 客户端从 Consul 获取服务实例列表
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	frontendUtils.MustHandleErr(err)

	// 使用对应的IDL客户端
	CartClient, err = cartservice.NewClient("cart", client.WithResolver(r))
	frontendUtils.MustHandleErr(err)
}

func initCheckoutClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	frontendUtils.MustHandleErr(err)
	opts = append(opts, client.WithResolver(r))
	CheckoutClient, err = checkoutservice.NewClient("checkout", opts...)
	frontendUtils.MustHandleErr(err)
}
