package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/wsq1220/rpc-grpc-etcd-consul-micro/grpc/demo2_stream/serverandclient_stream/message"
	"google.golang.org/grpc"
)

//订单服务实现
type OrderServiceImpl struct {
}

//实现grpc双向流模式
func (os *OrderServiceImpl) GetOrderInfos(stream message.OrderService_GetOrderInfosServer) error {

	for {
		orderRequest, err := stream.Recv()
		if err == io.EOF {
			fmt.Println(" 数据读取结束 ")
			return err
		}
		if err != nil {
			panic(err.Error())
			return err
		}

		fmt.Println(orderRequest.GetOrderId())
		orderMap := map[string]message.OrderInfo{
			"201907300001": message.OrderInfo{OrderId: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
			"201907310001": message.OrderInfo{OrderId: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
			"201907310002": message.OrderInfo{OrderId: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
		}

		result := orderMap[orderRequest.GetOrderId()]
		//发送数据
		err = stream.Send(&result)
		if err == io.EOF {
			fmt.Println(err)
			return err
		}
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

func main() {

	server := grpc.NewServer()
	//注册
	message.RegisterOrderServiceServer(server, new(OrderServiceImpl))

	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		panic(err.Error())
	}
	log.Printf("[grpc] server is listening: %d", 8088)
	server.Serve(lis)

}
