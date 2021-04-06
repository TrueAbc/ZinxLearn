package ziface

// 定义服务端的接口
type IServer interface {
	//启动
	Start()
	//运行
	Serve()

	Stop()
	// 停止
}
