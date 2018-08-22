package main

import (
	"flag"
	"github.com/smallnest/rpcx/client"
	"thresher/cp2/handler"
	"context"
	"log"
)

var (
	addr = flag.String("addr", "localhost:8972","server address")
)
func main() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr,"")
	xclient := client.NewXClient("a.fake.service",client.Failtry,client.RandomSelect,d,client.DefaultOption)
	defer xclient.Close()

	args := &handler.Args{
		A: 10,
		B: 20,
	}
	reply := &handler.Reply{}

	err := xclient.Call(context.Background(),"mul",args,reply)
	if err != nil {
		log.Fatalf("faild to call: %v", err)
	}
	log.Printf("%d * %d = %d", args.A, args.B, reply.C)
}
