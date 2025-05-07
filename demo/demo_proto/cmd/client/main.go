package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi/echoservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	// build a consul resolver with the consul client
	r, err := consul.NewConsulResolver("192.168.3.6:8500")

	if err != nil {
		log.Fatal(err)
	}

	c, err := echoservice.NewClient("demo_proto", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}

	for {
		res, err := c.Echo(context.TODO(), &pbapi.Request{Message: "hello"})
		if err != nil {
			// log.Fatal(err)
			continue
		}

		fmt.Println(res)

		time.Sleep(time.Duration(time.Second))
	}

}
