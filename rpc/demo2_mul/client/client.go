package main

import (
	"fmt"
	"net/rpc"

	"github.com/wsq1220/rpc-grpc-etcd-consul-micro/rpc/demo2_mul/param"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8082")
	if err != nil {
		panic(err.Error())
	}

	var res *float32
	param := param.Param{Arg1: 2.2, Arg2: 3.3}
	err = client.Call("MathCal.Cal", param, &res)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(*res)
}
