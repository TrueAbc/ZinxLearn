package znet

import (
	"fmt"
	"net"
	"trueabc.top/zinx/utils"
	"trueabc.top/zinx/ziface"
)

// 定义server模块
type Server struct {
	Name string

	IPVersion string

	IP string

	Port int64

	// 當前server添加的router
	Router ziface.IRouter
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router success!")
}

func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name: %s, listener at IP: %s, Port: %d is starting\n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	go func() {
		// 1. 获取一个Tcp地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp address err:", err)
			return
		}
		// 2. 监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, " err", err)
			return
		}
		fmt.Println("start Zinx Server ", s.Name)
		var cid uint32
		// 3.  等待链接
		for {
			// 客户端链接过来阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err: ", err)
				continue
			}
			// 將處理鏈接的方法和conn綁定, 得到鏈接模塊
			dealConn := NewConnection(conn, cid, s.Router)
			go dealConn.Start()

			cid++
		}
	}()
}

func (s *Server) Stop() {
	// TODO 将相关的资源进行回收
}

func (s *Server) Serve() {
	s.Start()

	// 阻塞状态
	select {}

}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}
	return s
}
