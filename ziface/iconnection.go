package ziface

import "net"

type IConnection interface {
	// 启动链接,
	Start()
	// 结束当前链接
	Stop()
	// 获取当前链接的socket conn
	GetTCPConnection() *net.TCPConn
	// 获取当前链接的ID
	GetConnID() uint32
	// 获取远程客户端的TCP的端口地址
	GetRemoteAddr() net.Addr
	//  发送数据
	SendMsg(msgId uint32, data []byte) error

	// 设置连接属性
	SetProperty(key string, value interface{})
	// 获取连接属性
	GetProperty(key string) (interface{}, error)
	// 移除连接属性
	RemoveProperty(key string)
}

// 定义处理链接的业务方法
type HandleFunc func(*net.TCPConn, []byte, int) error
