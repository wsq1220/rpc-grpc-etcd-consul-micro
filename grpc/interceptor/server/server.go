package main

import (
	"context"
	"fmt"
	"net"

	"github.com/wsq1220/rpc-grpc-etcd-consul-micro/grpc/interceptor/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type MathManager struct {
}

func (m *MathManager) AddMethod(ctx context.Context, request *message.RequestArgs) (response *message.Response, err error) {
	fmt.Println("server add method...")
	result := request.Args1 + request.Args2
	fmt.Printf("result: %v\n", result)

	response = &message.Response{
		Code:    0,
		Message: "success",
	}

	return
}

func main() {
	//TLS认证
	creds, err := credentials.NewServerTLSFromFile("./keys/server.pem", "./keys/server.key")
	if err != nil {
		grpclog.Fatal("加载在证书文件失败", err)
	}

	//实例化grpc server, 开启TLS认证
	server := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(TokenInterceptor))

	message.RegisterMathServiceServer(server, new(MathManager))

	lis, err := net.Listen("tcp", ":8093")
	if err != nil {
		panic(err.Error())
	}
	server.Serve(lis)
}

// 自定义拦截器实现
func TokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, exist := metadata.FromIncomingContext(ctx)
	if !exist {
		return nil, status.Errorf(codes.Unauthenticated, "无Token认证信息")
	}

	var appKey string
	var appSecret string

	if key, ok := md["appid"]; ok {
		appKey = key[0]
	}

	if secret, ok := md["appkey"]; ok {
		appSecret = secret[0]
	}

	if appKey != "hello" || appSecret != "20190812" {
		return nil, status.Errorf(codes.Unauthenticated, "Token 不合法")
	}

	// 通过token验证，继续处理请求
	return handler(ctx, req)
}
