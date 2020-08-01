package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"github.com/wsq1220/rpc-grpc-etcd-consul-micro/rpc/demo2_mul/param"

)

type MathCal struct {

}

func (m *MathCal) Cal(param param.Param, resp *float32) (err error) {
	*resp = param.Arg1 + param.Arg2
	return
}

func main() {
	mathCal := &MathCal{}
	if err := rpc.Register(mathCal); err != nil {
		panic(err.Error())
	}

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		panic(err.Error())
	}

	rpc.HandleHTTP()
	fmt.Printf("server is listening: %d\n", 8082)
	http.Serve(lis, nil)
}