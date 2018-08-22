package main

import (
	"flag"
	"github.com/smallnest/rpcx/client"
	"thresher/cp2/handler"
	"context"
	"log"
	"time"
)

var (
	addr1 = flag.String("addr1","tcp@localhost:8972","server1 address")
	addr2 = flag.String("addr2", "tcp@localhost:9981","")
)
func main() {
	flag.Parse()

	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1},{Key: *addr2}})

	xclient := client.NewXClient("Arith",client.Failover,client.RoundRobin,d,client.DefaultOption)
	defer xclient.Close()

	args := &handler.Args{
		A: 10,
		B: 20,
	}
	for {
		reply := &handler.Reply{}
		err := xclient.Call(context.Background(),"Mul",args,reply)
		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}
		log.Printf("%d * %d = %d", args.A, args.B, reply.C)
		time.Sleep(1e9)
	}

}
