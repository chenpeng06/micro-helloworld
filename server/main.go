package main

import (
	"context"
	hellopb "github.com/chenpeng06/micro-helloworld/proto"
	"github.com/micro/go-micro/v2"
)

type HelloServer struct {
}

func (hs *HelloServer) SayHello(ctx context.Context, in *hellopb.SayRequest, out *hellopb.SayResponse) error {
	out.Answer = "我们的口号是： " + in.Message + "。"
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("cap.imooc.server"),
	)
	// 初始化
	service.Init()
	hellopb.RegisterHelloHandler(service.Server(), new(HelloServer))

	if err := service.Run(); err != nil {
		panic(err)
	}
}
