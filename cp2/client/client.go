package main

import (
	"flag"
	"thresher/cp2/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	flag.Parse()
	r := gin.Default()

	r.GET("/ping", handler.Get)
	r.GET("/hello")
	r.Run() // 在 0.0.0.0:8080 上监听并服务

}
