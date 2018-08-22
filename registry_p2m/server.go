package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"thresher/cp2/handler"
)

var (
	addr1 = flag.String("addr1","localhost:8972","server1 address")
	addr2 = flag.String("addr2","localhost:9981", "server2 address")
)
func main() {
	flag.Parse()
	go createServer(*addr1)
	go createServer(*addr2)
	select {}
}
func createServer(addr string) {
	s := server.NewServer()
	s.RegisterName("Arith", new(handler.Arith),"")
	s.Serve("tcp", addr)
}