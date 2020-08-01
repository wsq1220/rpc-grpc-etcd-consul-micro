package main

import (
	"fmt"
	"net/rpc"
	"time"

	message "github.com/wsq1220/rpc-grpc-etcd-consul-micro/rpc/demo3_protobuf/proto"
)

func main() {

	client, err := rpc.DialHTTP("tcp", "localhost:8083")
	if err != nil {
		panic(err.Error())
	}
	timeStamp := time.Now().Unix()
	request := message.OrderRequest{OrderId: "202005310003", TimeStamp: timeStamp}
	var response *message.OrderInfo
	err = client.Call("OrderService.GetOrderInfo", request, &response)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(*response)
}
