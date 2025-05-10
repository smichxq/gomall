package rpc

import (
	"log"
	"sync"

	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient userservice.Client

	once sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
	// 客户端从 Consul 获取服务实例列表
	r, err := consul.NewConsulResolver("192.168.3.6:8500")
	if err != nil {
		log.Fatal("NewConsulRegister", err)
		return
	}

	// 使用对应的IDL客户端
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}
