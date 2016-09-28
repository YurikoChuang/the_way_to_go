package rpc_objects

import (
	//"net"
)

type Args struct{
	N,M int
}

func (t *Args)Multiply(args *Args , reply *int) error{  //不要返回net.Error //net.Error {
	*reply = args.N * args.M
	return nil
}