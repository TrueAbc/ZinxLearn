package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"trueabc.top/zinx/utils"
	"trueabc.top/zinx/ziface"
)

/*
链接模块
*/

type Connection struct {
	// 当前connection隶属的server
	TcpServer ziface.IServer

	Conn *net.TCPConn

	ConnID uint32

	// 当前链接状态
	isClosed bool

	// 当前链接的业务方法
	handler ziface.HandleFunc

	// 无缓冲管道, 用于读写Goroutine之间的消息通信
	msgChan chan []byte

	// 告知当前链接退出的chan, reader告知writer退出
	ExitChan chan bool

	// 消息的管理MsgID和对应的处理业务的API
	MsgHandler ziface.IMsgHandle
}

/*
 写消息的Goroutine, 专门发送消息
*/
func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running]")
	defer fmt.Println(c.GetRemoteAddr().String(), "[conn writer exit!]")

	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data err ", err)
				return
			}
		case <-c.ExitChan:
			// 代表reader退出, writer也退出
			return

		}
	}
}

// 链接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println("connId = ", c.ConnID, "Reader is exit, remote addr is  ", c.GetRemoteAddr().String())
	defer c.Stop()

	for true {
		// 读取客户端的数据到buf中
		//buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		//cnt, err := c.Conn.Read(buf)
		//if err != nil {
		//	fmt.Println("recv buf err: ", err)
		//	continue
		//}
		// 创建一个拆包对象
		dp := NewDataPack()
		// 获取客户端的Msg Header
		headerData := make([]byte, dp.GetHeaderLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headerData); err != nil {
			fmt.Println("read msg error ", err)
			break
		}
		// 拆包, 得到MsgId和datalen
		msg, err := dp.Unpack(headerData)
		if err != nil {
			fmt.Println("unpack err: ", err)
			break
		}

		// 根据dataLen读取data,
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error: ", err)
				break
			}
		}

		msg.SetMsg(data)

		// 得到當前的request
		req := Request{
			conn: c,
			msg:  msg,
		}
		if utils.GlobalObject.WorkerPoolSize > 0 {
			// 调用当前的处理逻辑, 路由的方法
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			go c.MsgHandler.DoMsgHandler(&req)
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start()...ConnID= ", c.ConnID)

	// 启动当前的读取业务
	go c.StartReader()
	go c.StartWriter()

}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop()...ConnID= ", c.ConnID)

	if c.isClosed == true {
		return
	}
	c.isClosed = true

	// 告知writer關閉
	c.ExitChan <- true

	c.TcpServer.GetConnMgr().Remove(c)
	// 资源回收
	c.Conn.Close()

	close(c.ExitChan)
}

func (c Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 提供一个SendMsg方法, 将数据先进行封包, 再发送給writer的chan
func (c Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("Connection closed when send msg")
	}

	// 将data进行封包
	dp := NewDataPack()
	msg := NewMessage(msgId, data)
	msgStream, err := dp.Pack(msg)
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("Pack error msg")
	}

	c.msgChan <- msgStream

	return nil
}

func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, handle ziface.IMsgHandle) *Connection {
	c := &Connection{
		TcpServer:  server,
		Conn:       conn,
		ConnID:     connID,
		MsgHandler: handle,
		msgChan:    make(chan []byte, 0),
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
	}
	c.TcpServer.GetConnMgr().Add(c)
	// 将connection加入到manager中
	return c
}
