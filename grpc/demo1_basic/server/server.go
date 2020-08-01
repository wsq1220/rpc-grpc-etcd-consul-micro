package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	message "github.com/wsq1220/rpc-grpc-etcd-consul-micro/grpc/demo1_basic/proto"
	"google.golang.org/grpc"
)

type OrderServiceImpl struct {
}

// OrderServiceServer interface
//GetOrderInfo(context.Context, *OrderRequest) (*OrderInfo, error)
func (os *OrderServiceImpl) GetOrderInfo(ctx context.Context, request *message.OrderRequest) (*message.OrderInfo, error) {

	orderMap := map[string]message.OrderInfo{
		"202005310001": message.OrderInfo{OrderId: "202005310001", OrderName: "衣服", OrderStatus: "已付款"},
		"202005310002": message.OrderInfo{OrderId: "202005310002", OrderName: "零食", OrderStatus: "已付款"},
		"202005310003": message.OrderInfo{OrderId: "202005310003", OrderName: "食品", OrderStatus: "未付款"},
	}

	var response *message.OrderInfo
	current := time.Now().Unix()
	if request.TimeStamp > current {
		*response = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
	} else {
		result := orderMap[request.OrderId]
		if result.OrderId != "" {
			fmt.Println(result)
			return &result, nil
		} else {
			return nil, errors.New("server error")
		}
	}

	return response, nil
}

func main() {
	// grpc.Server
	server := grpc.NewServer()
	// func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer)
	message.RegisterOrderServiceServer(server, &OrderServiceImpl{})

	lis, err := net.Listen("tcp", ":8085")
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("server is listening: %d\n", 8085)
	server.Serve(lis)
}
