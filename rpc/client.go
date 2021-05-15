package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 与服务端创建连接
	client,err := rpc.DialHTTP("tcp","127.0.0.1:8082")
	if err != nil {
		fmt.Println(err)
		fmt.Println("连接失败")
	}
	// 返回值
	var name string

	//  远程调用函数:被调用的方法，传入值 ，返回值
	err = client.Call("User.GetStr","zhiliao",&name)
	if err != nil{
		fmt.Println("call出错")
	}
	fmt.Println("结果为：",name)
}