// go run -tags etcd server.go
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr     = flag.String("addr", "132.232.109.253:8973", "server address") // 服务调用地址
	addr2 = flag.String("addr2","0.0.0.0:8973","server addr") // 服务运行地址
	etcdAddr = flag.String("etcdAddr", "132.232.109.253:2379", "etcd address")
	basePath = flag.String("base", "/rpcx", "prefix path")
)

type Arith2 int

func (t *Arith2) Mul(ctx context.Context, args *example.Args, reply *example.Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s)

	s.RegisterName("Arith", new(Arith2), "")
	s.Serve("tcp", *addr2)
}

func addRegistryPlugin(s *server.Server) {

	r := &serverplugin.EtcdRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		EtcdServers:    []string{*etcdAddr},
		BasePath:       *basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Second * 10,
	}

	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
