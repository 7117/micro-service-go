package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	hello "micro_service/grpc_proto/proto"
)

func main() {

	//方法一：遍历尝试连接集群的成员  解决寻找好的服务节点
	// WithInsecure：跳过证书的验证
	//conn, err := grpc.Dial("192.168.0.104:8080", grpc.WithInsecure())
	//defer recover()
	//arr := []string{"192.168.0.104:8080", "192.168.0.105:8080", "192.168.0.106:8080", "192.168.0.108:8080"}
	//for i := 0; i < len(arr); i++ {
	//	conn, err = grpc.Dial(arr[i], grpc.WithInsecure())
	//}

	//方法二：服务发现   解决寻找好的服务节点
	var waitIndex uint64

	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client_api, _ := api.NewClient(config)
	services, _, _ := client_api.Health().Service("my_say_hi", "say_hi", true, &api.QueryOptions{
		WaitIndex: waitIndex,
	})

	// 遍历出返回的切片 更精细的选择
	//for _,service := range services {
	//	fmt.Println(service.Service.Address)
	//	fmt.Println(service.Service.Port)
	//}

	address := services[0].Service.Address
	port := services[0].Service.Port

	url := fmt.Sprintf("%s:%d", address, port)
	fmt.Println(url)
	//http.Get("http://127.0.0.1:8500/v1/health/service/my_say_hi?passing=true")

	//拨号服务端
	conn, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("连接错误")
	}

	//连接
	client := hello.NewTestServiceClient(conn)

	//调用服务端的方法
	rep, err1 := client.SayHi(context.Background(), &hello.HiRequest{Name: "zhiliao"})
	if err1 != nil {
		fmt.Println("返回出错")
	}

	fmt.Println(rep.Ret)
}
