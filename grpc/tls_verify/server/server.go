package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/wsq1220/rpc-grpc-etcd-consul-micro/grpc/tls_verify/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

// type MathServiceServer interface {
// 	// 服务
// 	AddMethod(context.Context, *RequestArgs) (*Response, error)
// }

type MathManage struct {
}

func (m *MathManage) AddMethod(ctx context.Context, req *message.RequestArgs) (response *message.Response, err error) {
	fmt.Println("server method...")
	result := req.Args1 + req.Args2
	fmt.Println("result: ", result)

	response = &message.Response{
		Code:    0,
		Message: "success",
	}

	return
}

func main() {

	// grpc.WithPerRPCCredentials()

	//TLS认证
	creds, err := credentials.NewServerTLSFromFile("../keys/server.pem", "./keys/server.key")
	if err != nil {
		grpclog.Fatal("load cer file failed, err: ", err)
	}

	//实例化grpc server, 开启TLS认证
	server := grpc.NewServer(grpc.Creds(creds))

	message.RegisterMathServiceServer(server, &MathManage{})

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		panic(err.Error())
	}

	log.Printf("[grpc] server is listening: %d", 8089)
	server.Serve(lis)
}
