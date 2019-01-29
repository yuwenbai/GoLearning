package main

import (
	"context"
	"fmt"

	proto "../helloworld/proto"

	micro "github.com/micro/go-micro"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
)

func main() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(micro.Registry(reg), micro.Name("greeter.client"))
	service.Init()
	greeter := proto.NewGreeterService("helloworld", service.Client())
	rsp, err := greeter.SayHello(context.TODO(), &proto.HelloRequest{Name: "John, how are you?"})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.Message)
}
