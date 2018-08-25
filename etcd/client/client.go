package main

import (
	"github.com/gin-gonic/gin"
	"thresher/etcd/handler"
)

func main() {
	g := gin.Default()
	g.GET("/ping",handler.Get)
	g.Run()
}
