package main

import (
	"fmt"
	"trueabc.top/zinx/ziface"
	"trueabc.top/zinx/znet"
)

// ping test 自定義路由
type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) Handler(request ziface.IRequest) {
	fmt.Println("Call Router handler")
	// 先读取客户端的数据, 再ping... ping... ping...
	fmt.Printf("recv from client: MsgId= %d, data=%s", request.GetMsgID(), request.GetData())

	err := request.GetConnection().SendMsg(200, []byte("ping router is working"))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloRouter struct {
	znet.BaseRouter
}

func (pr *HelloRouter) Handler(request ziface.IRequest) {
	fmt.Println("Call Hello handler")
	// 先读取客户端的数据, 再ping... ping... ping...
	fmt.Printf("recv from client: MsgId= %d, data=%s", request.GetMsgID(), request.GetData())

	err := request.GetConnection().SendMsg(201, []byte("hello hello hello zinx"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServer("[zinx 0.8]")

	s.SetOnConnStart(func(connection ziface.IConnection) {
		fmt.Println("---------------> Start deal with client connection")
		connection.SendMsg(202, []byte("DoConnection end"))
	})

	s.SetOnConnStop(func(connection ziface.IConnection) {
		fmt.Println("close connection with client <-------------------------")
		connection.SendMsg(202, []byte("DoConnection end"))
	})

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.Serve()
}
