package handler

import (
	"context"
	"fmt"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith struct {

}

func (a *Arith) Mul(ctx context.Context, args Args,reply *Reply) error {
	reply.C = args.A * args.B
	fmt.Printf("call: %d * %d = %d\n", args.A, args.B, reply.C)
	return nil
}