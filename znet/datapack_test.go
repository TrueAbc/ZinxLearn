package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)

// 负责测试datapack拆包 封包的单元测试
func TestDataPack(t *testing.T) {

	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("server accept err: ", err)
			}

			go func(conn2 net.Conn) {
				// 拆包
				//1. 获取header信息
				dp := NewDataPack()
				for true {
					headData := make([]byte, dp.GetHeaderLen())
					_, err := io.ReadFull(conn2, headData)
					if err != nil {
						fmt.Println("read head error")
						return
					}
					msgHeader, _ := dp.Unpack(headData)

					if msgHeader.GetMsgLen() > 0 {
						// msg 有数据, 读取内容
						msg := msgHeader.(*Message)
						msg.Data = make([]byte, msg.GetMsgLen())

						// 根据dataLen长度再次从io中读取
						io.ReadFull(conn2, msg.Data)

						fmt.Println("---->Recv MsgID: ", msg.Id, "datalen= ", msg.DataLen, " data= ", string(msg.Data))
					}
				}
				//2. 获取内容信息

			}(conn)
		}
	}()

	// 模拟客户端
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	dp := NewDataPack()

	// 模拟粘包过程

	// 第一个包, 第二个包, 粘在一起
	msg1 := &Message{
		Id:      1,
		DataLen: 5,
		Data: []byte(
			"abcde"),
	}
	msg2 := &Message{
		Id:      2,
		DataLen: 10,
		Data: []byte(
			"hello world !"),
	}
	S1, _ := dp.Pack(msg1)
	S2, _ := dp.Pack(msg2)

	sendData1 := append(S1, S2...)
	conn.Write(sendData1)

	time.Sleep(time.Second * 5)

}
