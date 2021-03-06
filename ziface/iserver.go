package ziface

// 定义服务端的接口
type IServer interface {
	//启动
	Start()
	//运行
	Serve()

	Stop()
	// 停止

	// 路由功能, 給當前的服務注冊一個路由功能，供給客戶端鏈接使用
	AddRouter(msgId uint32, router IRouter)

	GetConnMgr() IConnManager

	SetOnConnStart(func(connection IConnection))
	SetOnConnStop(func(connection IConnection))

	CallOnConnStart(connection IConnection)
	CallOnConnStop(connection IConnection)
}
