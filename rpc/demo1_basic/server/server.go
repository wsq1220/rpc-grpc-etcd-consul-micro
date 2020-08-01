package main

import (
	"math"
	"net"
	"net/http"
	"net/rpc"

	"github.com/emicklei/go-restful/log"
)

type MathCal struct {
}

func (m *MathCal) Cal(req float32, resp *float32) (err error) {
	*resp = math.Pi * req * req
	return
}

func main() {
	// 初始化指针数据类型
	mathCal := &MathCal{}
	// 将服务对象进行注册
	// rpc.Register(mathCal)
	if err := rpc.RegisterName("MathCal", mathCal); err != nil {
		log.Printf("register rpc ser failed, err: %v\n", err)
		panic(err.Error())
	}

	// 把mathCal中提供的服务注册到http,方便调用者可以利用http的方式进行数据传递
	rpc.HandleHTTP()

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err.Error())
	}

	log.Printf("server is listening: %d", 8081)
	http.Serve(lis, nil)
}
