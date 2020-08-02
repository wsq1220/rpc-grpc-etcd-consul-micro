package main

import (
	"github.com/wsq1220/rpc-grpc-etcd-consul-micro/grpc/token/message"
	"google.golang.org/grpc"
)

type TokenAuthentication struct {
	AppKey string 
	AppSecret string
}

// type PerRPCCredentials interface {
// 	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
// 	RequireTransportSecurity() bool
// }

//组织token认证的metadata信息
func (ta *TokenAuthentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"AppKey": ta.AppKey, 
		"AppSecret": ta.AppSecret,
	}, nil
}

//是否基于TLS认证进行安全传输
func (ta *TokenAuthentication) RequireTransportSecurity() book {
	return true
}

func main() {
	// tls conn
	creds, err := credentials.NewClientTLSFromFile("../keys/server.pem", "go-grpc-example")
	if err != nil {
		panic(err.Error())
	}

	auth := TokenAuthentication {
		AppKey: "my_key",
		AppSecret: "my_secret",
	}

	conn, err := grpc.Dial("localhost:8089", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	serviceClient := message.NewMathServiceClient(conn)
	addArgs := message.RequestArgs{Args1: 3, Args2: 5}

	response, err := serviceClient.AddMethod(context.Background(), &addArgs)
	if err != nil {
		grpclog.Fatal(err.Error())
	}
	fmt.Println(response.GetCode(), response.GetMessage())
}

