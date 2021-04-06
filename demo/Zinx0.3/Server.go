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

func (pr *PingRouter) PreHandler(request ziface.IRequest) {
	fmt.Println("Call Router pre handler")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping..\n"))
	if err != nil {
		fmt.Println("call back before ping error")
	}

}

func (pr *PingRouter) Handler(request ziface.IRequest) {
	fmt.Println("Call Router handler")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping... ping..\n"))
	if err != nil {
		fmt.Println("call back  ping error")
	}
}

func (pr *PingRouter) PostHandler(request ziface.IRequest) {
	fmt.Println("Call Router post handler")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping..\n"))
	if err != nil {
		fmt.Println("call back after ping error")
	}
}

func main() {
	s := znet.NewServer("[zinx 0.3]")

	s.AddRouter(&PingRouter{})
	s.Serve()
}
