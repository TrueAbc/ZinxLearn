package znet

import (
	"fmt"
	"trueabc.top/zinx/ziface"
)

/*
 消息处理模块的实现
*/

type MsgHandler struct {
	// 存放每个MsgID所对应的处理方法
	Apis map[uint32]ziface.IRouter
}

// 初始化/ 创建MsgHandler
func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

// 调度/ 执行对应的Router消息处理方法
func (m MsgHandler) DoMsgHandler(request ziface.IRequest) {
	// 从request中找到msgID
	// 根据id处理业务
	handler, ok := m.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgId= ", request.GetMsgID(), " is not found!")
	}
	handler.PreHandler(request)
	handler.Handler(request)
	handler.PostHandler(request)

	return
}

func (m *MsgHandler) AddRouter(msgId uint32, router ziface.IRouter) {
	// 判断当前的id如果已经绑定了就不再处理
	if _, ok := m.Apis[msgId]; ok {
		// id已经注册
		fmt.Println("msgId ", msgId, " routerHandler has been registered ")
		return
	}

	m.Apis[msgId] = router
	fmt.Println("add api msgId= ", msgId, "success!")

}
