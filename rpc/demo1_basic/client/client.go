package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8081")
	if err != nil {
		panic(err.Error())
	}

	var req float32
	req = 6

	// var resp *float32
	// err = client.Call("MathCal.Cal", req, &resp)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// 异步
	var syncResp *float32
	syncCall := client.Go("MathCal.Cal", req, &syncResp, nil)
	// fmt.Println(*syncResp)  会报错
	done := <-syncCall.Done
	fmt.Println(done)
	fmt.Println(*syncResp)
}
