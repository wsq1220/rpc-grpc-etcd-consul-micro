package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/wsq1220/rpc-grpc-etcd-consul-micro/grpc/demo2_stream/server_stream/message"
	"google.golang.org/grpc"
)

//订单服务实现
type OrderServiceImpl struct {
}

//获取订单信息s
func (os *OrderServiceImpl) GetOrderInfos(request *message.OrderRequest, stream message.OrderService_GetOrderInfosServer) error {
	fmt.Println(" 服务端流 RPC 模式")

	orderMap := map[string]message.OrderInfo{
		"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
		"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
		"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
	}
	for id, info := range orderMap {
		if time.Now().Unix() >= request.TimeStamp {
			fmt.Println("订单序列号ID：", id)
			fmt.Println("订单详情：", info)
			//通过流模式发送给客户端
			stream.Send(&info)
		}
	}
	return nil
}

func main() {
	server := grpc.NewServer()
	message.RegisterOrderServiceServer(server, &OrderServiceImpl{})

	lis, err := net.Listen("tcp", ":8087")
	if err != nil {
		panic(err.Error())
	}

	log.Printf("[grpc] server is listening: %d", 8087)
	server.Serve(lis)
}
