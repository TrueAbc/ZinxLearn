package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"trueabc.top/zinx/utils"
	"trueabc.top/zinx/ziface"
)

/*
 封包, 拆包的模块
*/

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (d DataPack) GetHeaderLen() uint32 {
	// datalen uint32 4byte, dataid uint32 4byte
	return utils.GlobalObject.MsgHeaderLen
}

func (d DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 创建一个存放字节流的缓冲
	databuf := bytes.NewBuffer([]byte{})
	// 将dataLen写入dataBuf
	if err := binary.Write(databuf, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}

	// 将msgId写入dataBuf
	if err := binary.Write(databuf, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	if err := binary.Write(databuf, binary.LittleEndian, msg.GetMsg()); err != nil {
		return nil, err
	}
	return databuf.Bytes(), nil
}

// 拆包方法, 将head信息读出, 再跟进head信息的data的长度, 读取相关的内容信息
func (d DataPack) Unpack(data []byte) (ziface.IMessage, error) {
	// 创建一个二进制数据的reader
	dataBuf := bytes.NewReader(data)

	msg := &Message{}
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	// 判断是否超过最大的包长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too Large msg data recv!")
	}

	return msg, nil
}
