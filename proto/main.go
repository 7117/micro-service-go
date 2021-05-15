package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	hello "micro_service/proto/protofile"
)

func main() {
	name := "aaaa"
	helloData := &hello.Hello{Name: &name, Age: 19, Addr: "vvvv"}
	fmt.Println(helloData)
	fmt.Println(helloData.Age)
	fmt.Println(helloData.GetName())
	fmt.Println(*helloData.Name)

	//	编码
	byteData, _ := proto.Marshal(helloData)
	fmt.Println(byteData)

	//	解码
	hello2Data := &hello.Hello{}
	proto.Unmarshal(byteData, hello2Data)
	fmt.Println(hello2Data)
	fmt.Println(hello2Data.GetName())
}
