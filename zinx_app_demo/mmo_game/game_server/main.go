package main

import "trueabc.top/zinx/znet"

func main() {
	s := znet.NewServer("MMO Game Server")

	// 注册路由业务

	// 连接创建和销毁的钩子函数

	s.Serve()
}
