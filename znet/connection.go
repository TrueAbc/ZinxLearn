package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"trueabc.top/zinx/ziface"
)

/*
链接模块
*/

type Connection struct {
	Conn *net.TCPConn

	ConnID uint32

	// 当前链接状态
	isClosed bool

	// 当前链接的业务方法
	handler ziface.HandleFunc

	// 告知当前链接退出的chan
	ExitChan chan bool

	// 該鏈接處理的router方法
	Router ziface.IRouter
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

		// 调用当前的处理逻辑, 路由的方法
		go func(request ziface.IRequest) {
			c.Router.PreHandler(request)
			c.Router.Handler(request)
			c.Router.PostHandler(request)
		}(&req)
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start()...ConnID= ", c.ConnID)

	// 启动当前的读取业务
	go c.StartReader()

	//TODO 启动写数据的业务
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop()...ConnID= ", c.ConnID)

	if c.isClosed == true {
		return
	}
	c.isClosed = true

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

// 提供一个SendMsg方法, 将数据先进行封包, 再发送
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
	if _, err := c.Conn.Write(msgStream); err != nil {
		fmt.Println("Write msg id: ", msgId, " error: ", err)
		return errors.New("conn write error")
	}

	return nil
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}
