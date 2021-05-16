package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello "micro_service/grpc_proto/proto"
	"net"
)

type Server struct {}

func (s *Server) SayHi(ctx context.Context, in *hello.HiRequest) (*hello.HiResponse, error)  {

	rep := hello.HiResponse{Ret:"hi, "+ in.Name}
	return &rep,nil

}

func main()  {
	// 创建grpc服务
	server := grpc.NewServer()

	hello.RegisterTestServiceServer(server,&Server{})

	listen,err := net.Listen("tcp","127.0.0.1:8088")

	fmt.Println("监听8088端口...")
	if err != nil {
		fmt.Println("网络错误")
	}

	server.Serve(listen)   // grpc服务启动
}