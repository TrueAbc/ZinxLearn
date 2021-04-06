package znet

import (
	"fmt"
	"net"
	"trueabc.top/zinx/ziface"
)

// 定义server模块
type Server struct {
	Name string

	IPVersion string

	IP string

	Port int64
}

func (s *Server) Start() {
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
		// 3.  等待链接
		for {
			// 客户端链接过来阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err: ", err)
				continue
			}

			go func(conn net.Conn) {
				for true {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err")
						continue
					}

					// 回显
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back err: ", err)
						continue
					}
				}
			}(conn)
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
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8889,
	}
	return s
}
