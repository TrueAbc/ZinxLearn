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

	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping..."))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServer("[zinx 0.5]")

	s.AddRouter(&PingRouter{})
	s.Serve()
}
