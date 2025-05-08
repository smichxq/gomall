package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi/echoservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
)

func main() {
	// build a consul resolver with the consul client
	// r, err := consul.NewConsulResolver("192.168.3.6:8500")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// cli, err := echoservice.NewClient("demo_proto", client.WithResolver(r))

	// 单个微服务传递
	cli, err := echoservice.NewClient(
		"demo_proto",
		client.WithHostPorts("localhost:8888"),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	if err != nil {
		klog.Fatal(err)
	}

	// 适用于多个微服务传递
	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto_client")

	res, err := cli.Echo(ctx, &pbapi.Request{Message: "error"})
	var bizErr *kerrors.GRPCBizStatusError

	if err != nil {
		// 将err映射到bizErr
		ok := errors.As(err, &bizErr)
		if ok {
			fmt.Println("%#v", bizErr)
		}
		log.Fatal(err)
	}

	// res, err := cli.Echo(context.TODO(), &pbapi.Request{Message: "hello"})
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println(res)
}
