package main

import (
        "github.com/micro/go-log"
	"net/http"

        "github.com/micro/go-web"
        "micro_web_test/handler"
)

func main() {
	// create new web service
        service := web.NewService(
                //go.micro是命名空间
                //web是type
                //alias后面紧跟是别名
                //fqdn   全限定域名   就是全称  全部自己定义
                web.Name("go.micro.web.micro_web_test"),
                web.Version("latest"),
        )

	// initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/example/call", handler.ExampleCall)

	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
