package ziface

/*
	TCP协议封包和拆包的模块
*/

type IDataPack interface {
	GetHeaderLen() uint32

	Pack(msg IMessage) ([]byte, error)

	Unpack([]byte) (IMessage, error)
}
