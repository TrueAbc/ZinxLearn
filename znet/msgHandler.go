package znet

import (
	"fmt"
	"trueabc.top/zinx/utils"
	"trueabc.top/zinx/ziface"
)

/*
 消息处理模块的实现
*/

type MsgHandler struct {
	// 存放每个MsgID所对应的处理方法
	Apis map[uint32]ziface.IRouter

	// 负责worker获取任务的队列
	TaskQueue []chan ziface.IRequest

	WorkerPoolSize uint32
	// worker工作池的数量

}

// 初始化/ 创建MsgHandler
func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize, // 从全局配置中获取,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
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

// 启动一个worker工作池
func (m *MsgHandler) StartWorkerPool() {
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		// 启动一个worker, 开辟一个消息队列
		m.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		// 启动当前的worker, 阻塞等待channel传递消息过来
		go m.startOneWorker(i, m.TaskQueue[i])
	}
}

// 启动一个worker工作流程
func (m *MsgHandler) startOneWorker(workerId int, taskQueue chan ziface.IRequest) {
	fmt.Println("Worker ID=", workerId, " is started...")
	for true {
		select {
		// 如果有消息过来, 出列的就是一个客户端的request
		case req := <-taskQueue:
			m.DoMsgHandler(req)
		}
	}
}

// 将消息交给taskQueue, 由worker进行处理
func (m *MsgHandler) SendMsgToTaskQueue(request ziface.IRequest) {
	// 1. 将消息平均分配给worker
	workerID := request.GetConnection().GetConnID() % m.WorkerPoolSize
	fmt.Println("Add ConnID = ", request.GetConnection().GetConnID(),
		"request MsgID = ", request.GetMsgID(), " to WorkerID = ", workerID)
	// 2.  将消息发送给对应的worker的TaskQueue
	m.TaskQueue[workerID] <- request
}
