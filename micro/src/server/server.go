package main

import (
	"context"
	"fmt"

	proto "../helloworld/proto"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
)

type Greeter struct{}

func (s *Greeter) SayHello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloReply) error {
	rsp.Message = "Hello " + req.Name
	return nil
}

func main() {
	// 修改consul地址，如果是本机，这段代码和后面的那行使用代码都是可以不用的
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("helloworld"),
	)

	service.Init()
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
