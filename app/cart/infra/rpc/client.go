package rpc

import (
	"sync"

	"github.com/cloudwego/gomall/app/cart/conf"
	cartutils "github.com/cloudwego/gomall/app/cart/utils"
	"github.com/cloudwego/gomall/common/clientsuite"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	ProductClient productcatalogservice.Client

	once sync.Once

	ServiceName  = conf.GetConf().Kitex.Service
	RegisterAddr = conf.GetConf().Registry.RegistryAddress[0]
	err          error
)

func InitClient() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegisterAddr,
		}),
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartutils.MustHandleError(err)
}

func InitClientUnitTest(registryAddr string) {
	var opts []client.Option
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartutils.MustHandleError(err)
}
