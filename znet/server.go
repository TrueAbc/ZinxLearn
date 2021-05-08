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

	// 当前Server的消息管理模块, 用来绑定MsgID和对应的处理业务
	MsgHandler ziface.IMsgHandle

	// 当前server的连接管理模块, 每次与客户端建立连接后加入连接. 每次与客户端连接断开后删除连接
	// 添加连接之前判断当前的连接数量是否超过最大值
	ConnMgr ziface.IConnManager

	// 连接创建之后的钩子函数
	OnConnStart func(conn ziface.IConnection)

	OnConnStop func(conn ziface.IConnection)
}

func (s *Server) SetOnConnStart(f func(connection ziface.IConnection)) {
	s.OnConnStart = f
}

func (s *Server) SetOnConnStop(f func(connection ziface.IConnection)) {
	s.OnConnStop = f
}

func (s *Server) CallOnConnStart(connection ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("---------> Call OnConnStart() ...")
		s.OnConnStart(connection)
	}
}

func (s *Server) CallOnConnStop(connection ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---------> Call OnConnStop()....")
		s.OnConnStop(connection)
	}
}

func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
	fmt.Println("Add Router success!")
}

func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name: %s, listener at IP: %s, Port: %d is starting\n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	go func() {
		// 开启消息队列的工作池
		s.MsgHandler.StartWorkerPool()

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

			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				// todo 给客户端添加一个超过最大连接数的错误信息
				fmt.Println("==========》 Too many connection maxConn = ", utils.GlobalObject.MaxConn)
				conn.Close()
				continue
			}

			if err != nil {
				fmt.Println("Accept err: ", err)
				continue
			}
			// 將處理鏈接的方法和conn綁定, 得到鏈接模塊
			dealConn := NewConnection(s, conn, cid, s.MsgHandler)
			go dealConn.Start()

			cid++
		}
	}()
}

func (s *Server) Stop() {
	// TODO 将相关的资源进行回收
	fmt.Println("[STOP] Zinx Server name: ", s.Name)
	s.ConnMgr.ClearConn()

}

func (s *Server) Serve() {
	s.Start()

	// 阻塞状态
	select {}

}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandler(),
		ConnMgr:    NewConnManager(),
	}
	return s
}
