package znet

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

func (m Message) GetMsgId() uint32 {
	return m.Id
}

func (m Message) GetMsgLen() uint32 {
	return m.DataLen
}

func (m Message) GetMsg() []byte {
	return m.Data
}

func (m *Message) SetMsgId(u uint32) {
	m.Id = u
}

func (m *Message) SetMsgLen(u uint32) {
	m.DataLen = u
}

func (m *Message) SetMsg(data []byte) {
	m.Data = data
}

func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		id,
		uint32(len(data)),
		data,
	}
}
