// go run -tags consul server.go
package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	metrics "github.com/rcrowley/go-metrics"
	example "github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr       = flag.String("addr", "132.232.109.253:8972", "server address")
	addr2       = flag.String("addr2", "0.0.0.0:8972", "server address")
	consulAddr = flag.String("consulAddr", "127.0.0.1:8500", "consul address")
	basePath   = flag.String("base", "/rpcx", "prefix path")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s)

	s.RegisterName("Arith", new(example.Arith), "")
	err := s.Serve("tcp", *addr2)
	if err != nil {
		fmt.Println(err)
	}
}

func addRegistryPlugin(s *server.Server) {

	r := &serverplugin.ConsulRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		ConsulServers:  []string{*consulAddr},
		BasePath:       *basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}