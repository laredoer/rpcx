package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/smallnest/rpcx/client"
	"log"
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
	args := &handler.UserRequest{
		UserName: "admin",
		PassWord: "admin",
	}
	reply := &handler.Response{}
	err := xclient.Call(g, "Login", args, reply)
	if err != nil {
		log.Printf("failed to call: %v\n", err)
	}
	g.JSON(200,gin.H{
		"data":reply.Data,
	})
}