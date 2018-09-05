package model

import (
	"context"
	"flag"

	"github.com/rpcx-ecosystem/rpcx-examples3"
	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

type IUser struct {
	xclient client.XClient
}

var User *IUser

func init() {
	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	User = &IUser{
		xclient: client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption),
	}
}

func (u *IUser) Get(a, b int) (c int) {
	args := &example.Args{A: a, B: b}
	reply := &example.Reply{}
	err := u.xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		return reply.C
	}
	return reply.C
}
