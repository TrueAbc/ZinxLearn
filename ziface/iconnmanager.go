package ziface

// 链接管理模块的抽象

type IConnManager interface {

	// 添加
	Add(conn IConnection)
	// 删除
	Remove(conn IConnection)
	// 获取连接
	Get(connID uint32) (IConnection, error)
	// 得到当前连接总数
	Len() int
	// 清楚并终止所有连接
	ClearConn()
}
