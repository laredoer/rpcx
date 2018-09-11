// go run -tags consul server.go
package main

import (
	"flag"
	"fmt"
	"log"
	"thresher/srv/user-srv/handler"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr       = flag.String("addr", "localhost:8972", "server address")
	consulAddr = flag.String("consulAddr", "132.232.109.253:8500", "consul address")
	basePath   = flag.String("base", "/qu_video", "prefix path")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s)

	s.RegisterName("User", new(handler.User), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		fmt.Println(err)
	}
}

// 服务注册
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
