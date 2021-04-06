package znet

import (
	"errors"
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

// 定義當前客戶端綁定的handler, 之後應該由框架使用者進行自定義
func CallBack(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Connection Handler] CallBack to client")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err: ", err)
		return errors.New("CallBack to Client error")
	}

	return nil
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
			dealConn := NewConnection(conn, cid, CallBack)
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
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8889,
	}
	return s
}
