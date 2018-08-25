package model

import (
	"github.com/smallnest/rpcx/client"
	"flag"
	"github.com/gin-gonic/gin"
)
type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type IArith2 struct {
	xclient client.XClient
}

var (
	etcdAddr = flag.String("etcdAddr","localhost:2379","etcd address")
	basePath = flag.String("basePath","/rpcx","prefix path")
)

var Arith2 *IArith2

func init() {
	flag.Parse()
	d := client.NewEtcdDiscovery(*basePath,"Arith",[]string{*etcdAddr},nil)
	Arith2 = &IArith2{xclient : client.NewXClient("Arith",client.Failover,client.RoundRobin,d,client.DefaultOption)}

}

func (t *IArith2) Get(g *gin.Context,a *Args,r *Reply) (rep *Reply,err error)  {
	err = t.xclient.Call(g,"Mul",a,r)
	if err != nil {
		return nil,err
	}
	return r,nil
}


func (t *IArith2) Close () error{
	return  t.xclient.Close()
}
