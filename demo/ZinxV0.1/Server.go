package main

import "trueabc.top/zinx/znet"

func main() {
	//1. 创建server
	s := znet.NewServer("[zinx V0.1]")
	// 2. 启动Server
	s.Serve()
	// 3.
}
