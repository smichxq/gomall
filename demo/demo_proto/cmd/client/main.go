package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi/echoservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
)

func main() {
	// build a consul resolver with the consul client
	// r, err := consul.NewConsulResolver("192.168.3.6:8500")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// cli, err := echoservice.NewClient("demo_proto", client.WithResolver(r))
	cli, err := echoservice.NewClient(
		"demo_proto",
		client.WithHostPorts("localhost:8888"),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "demo_proto_client",
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	res, err := cli.Echo(context.TODO(), &pbapi.Request{Message: "hello"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
