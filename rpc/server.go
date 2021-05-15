package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type User struct {
}

func (n *User) GetStr(arg string, ret *string) error {
	*ret = "hello,"+arg + *ret
	return nil
}

func main() {
	// new 一个对象
	user := new(User)
	// 将对象注册到rpc服务中
	rpc.Register(user)
	// 把user中提供的服务注册到HTTP协议上,方便调用者可以利用http的方式进行数据传递
	rpc.HandleHTTP()

	// 监听指定端口
	listen, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil {
		fmt.Println("网络连接失败")
	}
	//service接受侦听器l上传入的HTTP连接，
	//http.Serve(listen,nil)
	fmt.Println("连接完成1")
	http.Serve(listen, nil)
	fmt.Println("连接完成2")
}
