package rpc

import (
	"sync"

	"github.com/cloudwego/gomall/app/checkout/conf"
	checkoutUtils "github.com/cloudwego/gomall/app/checkout/utils"
	"github.com/cloudwego/gomall/common/clientsuite"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	once          sync.Once
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	ServiceName   = conf.GetConf().Kitex.Service
	RegisterAddr  = conf.GetConf().Registry.RegistryAddress[0]
	err           error
	opts          = []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegisterAddr,
		}),
	}
)

func ClientInit() {
	once.Do(func() {
		cartClientInit()
		productClientInit()
		paymentClientInit()
	})
}

func paymentClientInit() {
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	checkoutUtils.MustHandleError(err)
}

func productClientInit() {
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	checkoutUtils.MustHandleError(err)
}

func cartClientInit() {
	CartClient, err = cartservice.NewClient("cart", opts...)
	checkoutUtils.MustHandleError(err)
}
