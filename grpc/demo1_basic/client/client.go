package main

import (
	"context"
	"fmt"
	"time"

	message "github.com/wsq1220/rpc-grpc-etcd-consul-micro/grpc/demo1_basic/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8085", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}

	// func NewOrderServiceClient(cc *grpc.ClientConn) OrderServiceClient
	orderServiceClient := message.NewOrderServiceClient(conn)

	orderRequest := &message.OrderRequest{OrderId: "202005310003", TimeStamp: time.Now().Unix()}
	// func (c *orderServiceClient) GetOrderInfo(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderInfo, error)
	orderInfo, err := orderServiceClient.GetOrderInfo(context.Background(), orderRequest)
	if orderInfo != nil {
		fmt.Println(orderInfo.GetOrderId())
		fmt.Println(orderInfo.GetOrderName())
		fmt.Println(orderInfo.GetOrderStatus())
	} else {
		fmt.Println(err.Error())
	}
}
