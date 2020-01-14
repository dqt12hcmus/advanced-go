package main

import (
	"context"
	"fmt"
	micro "github.com/micro/go-micro"
	proto "proto-server/greeter"
	"time"
)

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("1.0.1"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	service.Init()
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	counter++
	if counter > 7 && counter < 15 {
		time.Sleep(1000 * time.Millisecond)
	} else {
		time.Sleep(100 * time.Millisecond)
	}
	rsp.Greeting = fmt.Sprintf("Hello %s", req.Name)
	fmt.Println(counter)
	fmt.Printf("Responding with %s\n", rsp.Greeting)

	return nil
}

type Greeter struct {
}

var counter int
