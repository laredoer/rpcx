package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"thresher/cp2/handler"
)

var (
	addr = flag.String("addr","localhost:8972","server address")
)

func main () {
	flag.Parse()
	s := server.NewServer()
	s.RegisterName("Arith",new(handler.Arith),"")
	s.Serve("tcp",*addr)
}
