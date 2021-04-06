package znet

import (
	"fmt"
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
		// 读取客户端的数据到buf中, 最大512字节
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err: ", err)
			continue
		}

		// 得到當前的request
		req := Request{
			conn: c,
			data: buf[:cnt],
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

func (c Connection) Send(data []byte) error {
	panic("implement me")
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
