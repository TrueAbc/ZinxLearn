package znet

import (
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
}

func NewConnection(conn *net.TCPConn, connID uint32, callback ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		handler:  callback,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}
