package ziface

/*
IRequest 接口:
將客戶端請求的鏈接信息和請求的數據包裝在了一個Request中
*/

type IRequest interface {
	GetConnection() IConnection

	GetData() []byte

	GetMsgID() uint32
}
