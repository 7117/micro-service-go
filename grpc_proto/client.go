package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello "micro_service/grpc_proto/proto"
)

func main() {

	// WithInsecure：跳过证书的验证
	conn,err := grpc.Dial("127.0.0.1:8088",grpc.WithInsecure())

	if err != nil {
		fmt.Println("连接错误")
	}

	client := hello.NewTestServiceClient(conn)

	rep,err1 := client.SayHi(context.Background(),&hello.HiRequest{Name:"zhiliao"})
	if err1 != nil {
		fmt.Println("返回出错")
	}

	fmt.Println(rep.Ret)
}
