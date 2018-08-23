package main

import (
	"flag"
	"log"
	"thresher/cp2/handler"

	"github.com/gin-gonic/gin"

	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {

	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &handler.Args{
		A: 10,
		B: 20,
	}
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {

		reply := &handler.Reply{}
		err := xclient.Call(c, "Mul", args, reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}
		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		c.JSON(200, gin.H{
			"message": reply,
		})
	})
	r.Run() // 在 0.0.0.0:8080 上监听并服务

}
