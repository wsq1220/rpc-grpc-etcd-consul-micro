package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/wsq1220/rpc-grpc-etcd-consul-micro/grpc/demo2_stream/server_stream/message"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8087", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	orderServiceClient := message.NewOrderServiceClient(conn)
	request := message.OrderRequest{TimeStamp: time.Now().Unix()}
	orderInfosClient, err := orderServiceClient.GetOrderInfos(context.TODO(), &request)

	for {
		orderInfo, err := orderInfosClient.Recv()
		if err == io.EOF {
			fmt.Println("read data end")
			return
		}

		if err != nil {
			panic(err.Error())
		}
		fmt.Println("read data: ", orderInfo)
	}

}
