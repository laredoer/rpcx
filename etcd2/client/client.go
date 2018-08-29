package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/client"
)

var (
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd address")
	basePath = flag.String("base", "/rpcx", "prefix path")
	xclient  client.XClient
)

func init() {
	flag.Parse()

	d := client.NewEtcdDiscovery(*basePath, "Arith", []string{*etcdAddr}, nil)
	xclient = client.NewXClient("Arith", client.Failover, client.RoundRobin, d, client.DefaultOption)

}

func main() {
	defer xclient.Close()
	g := gin.Default()
	g.GET("/ping", Get)
	g.Run(":8888")
}

func Get(g *gin.Context) {
	args := &example.Args{
		A: 10,
		B: 20,
	}
	reply := &example.Reply{}
	err := xclient.Call(g, "Mul", args, reply)
	if err != nil {
		log.Printf("failed to call: %v\n", err)
	}
	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
