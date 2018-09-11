package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/smallnest/rpcx/client"
	"log"
	"net/http"
	"thresher/srv/user-srv/handler"
)

var (
	consulAddr = flag.String("consulAddr", "132.232.109.253:8500", "consul address")
	basePath   = flag.String("base", "/qu_video", "prefix path")
	xclient client.XClient
)

func init() {
	flag.Parse()

	d := client.NewConsulDiscovery(*basePath, "User", []string{*consulAddr}, nil)
	xclient = client.NewXClient("User", client.Failtry, client.RandomSelect, d, client.DefaultOption)
}



func main() {

	defer xclient.Close()

	g := gin.Default()
	g.GET("/ping",Get)
	g.Run(":8888")
}

func Get(g *gin.Context)  {
	var userLogin handler.UserRequest
	// 这个将通过 content-type 头去推断绑定器使用哪个依赖。
	if err := g.ShouldBind(&userLogin); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	reply := &handler.Response{}
	err := xclient.Call(g, "Login", &userLogin, reply)
	if err != nil || reply.Data == nil {
		g.JSON(http.StatusNotAcceptable,gin.H{
			"data":reply,
		})
		return
		log.Printf("failed to call: %v\n", err)
	}
	g.JSON(200,gin.H{
		"data":reply,
	})
}