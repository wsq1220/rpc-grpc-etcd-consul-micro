package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/wsq1220/rpc-grpc-etcd-consul-micro/grpc/demo2_stream/client_stream/message"
)

// type OrderServiceServer interface {
// 	AddOrderList(OrderService_AddOrderListServer) error
// }

//订单服务实现
type OrderServiceImpl struct {
}

//添加订单信息服务实现
func (os *OrderServiceImpl) AddOrderList(stream message.OrderService_AddOrderListServer) error {
	fmt.Println(" 客户端流 RPC 模式")

	for {
		//从流中读取数据信息
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println(" 读取数据结束 ")
			result := message.OrderInfo{OrderStatus: " 读取数据结束 "}
			return stream.SendAndClose(&result)
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		//打印接收到的数据
		fmt.Println(orderRequest.GetOrderId(), orderRequest.GetOrderName(), orderRequest.GetOrderStatus())
	}
}

func main() {
	server := grpc.NewServer()
	// 注册
	message.RegisterOrderServiceServer(server, &OrderServiceImpl{})

	lis, err := net.Listen("tcp", ":8086")
	if err != nil {
		panic(err.Error())
	}

	log.Printf("[grpc]server is listening on: %d", 8086)
	server.Serve(lis)
}
