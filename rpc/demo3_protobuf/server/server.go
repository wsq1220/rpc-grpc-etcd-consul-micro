package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"

	message "github.com/wsq1220/rpc-grpc-etcd-consul-micro/rpc/demo3_protobuf/proto"
)

type OrderService struct {
}

func (os *OrderService) GetOrderInfo(request message.OrderRequest, response *message.OrderInfo) error {

	orderMap := map[string]message.OrderInfo{
		"202005310001": message.OrderInfo{OrderId: "202005310001", OrderName: "衣服", OrderStatus: "已付款"},
		"202005310002": message.OrderInfo{OrderId: "202005310002", OrderName: "零食", OrderStatus: "已付款"},
		"202005310003": message.OrderInfo{OrderId: "202005310003", OrderName: "食品", OrderStatus: "未付款"},
	}

	current := time.Now().Unix()
	if request.TimeStamp > current {
		*response = message.OrderInfo{OrderId: "0", OrderName: "", OrderStatus: "订单信息异常"}
	} else {

		result := orderMap[request.OrderId]
		if result.OrderId != "" {
			*response = orderMap[request.OrderId]
		} else {
			return errors.New("server error")
		}
	}
	return nil
}

func main() {
	orderService := new(OrderService)

	rpc.Register(orderService)

	rpc.HandleHTTP()

	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("server is listening: %d\n", 8083)
	http.Serve(lis, nil)
}
