package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"trueabc.top/zinx/znet"
)

// 测试用客户端1
func main() {
	fmt.Println("client start...")
	// 1. 链接远程服务器
	time.Sleep(time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err: ", err)
		return
	}

	for true {
		// 发送封装的Message
		dp := znet.NewDataPack()
		msg, err := dp.Pack(znet.NewMessage(1, []byte("ZinxV0.6 client test 1 message\n")))
		if err != nil {

		}
		conn.Write(msg)

		// 从服务器接受消息
		binaryhead := make([]byte, dp.GetHeaderLen())
		if _, err := io.ReadFull(conn, binaryhead); err != nil {
			fmt.Println("read header error ", err)
			break
		}

		msgHead, err := dp.Unpack(binaryhead)
		if err != nil {
			fmt.Println("client unpack msgHead error ", err)
			break
		}

		if msgHead.GetMsgLen() > 0 {
			// msg是有数据的, 进行第二次数据读取
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data error ", err)
				return
			}
			fmt.Println("----> Recv Server Msg: ID= ", msg.Id, " len= ", msg.DataLen, " data= ", string(msg.Data))
		}
		time.Sleep(time.Second)
	}
}
