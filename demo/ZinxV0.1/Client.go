package main

import (
	"fmt"
	"net"
	"time"
)

// 测试用客户端1
func main() {
	fmt.Println("client start...")
	// 1. 链接远程服务器
	time.Sleep(time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("client start err: ", err)
		return
	}

	for true {
		_, err := conn.Write([]byte("hello Zinx"))
		if err != nil {
			fmt.Println("write to server err: ", err)
			continue
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err: ", err)
			continue
		}
		fmt.Println("server call back: ", string(buf[:cnt]), " cnt=", cnt)
		time.Sleep(time.Second)
	}
}
