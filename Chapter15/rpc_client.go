package main

import (
	"fmt"
	"log"
	"net/rpc"
	"./rpc_objects"
)
const serverAddress = "localhost"
func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")

	if err != nil {
		log.Fatal("error dialing: ", err)
	}

	//synchronous call
	args := &rpc_objects.Args{7,8}
	var reply int
	err = client.Call("Args.Multiply",args,&reply)  //现在client.Call()也是调用client.Go()
	if err != nil {
		log.Fatal("args error:" , err )
	}
	fmt.Printf("args: %d * %d = %d ", args.N,args.M,reply)

	//asynchronous call
	//call1 := client.Go("Args.Multiply",args,&reply,nil)
	//replyCall := <-call1.Done
}
