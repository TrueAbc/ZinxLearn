package znet

import "trueabc.top/zinx/ziface"

type Request struct {
	// 已經和客戶端建立好的鏈接
	conn ziface.IConnection

	// 客戶端請求的數據
	data []byte
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
