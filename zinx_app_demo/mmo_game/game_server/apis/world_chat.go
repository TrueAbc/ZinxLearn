package apis

import (
	"fmt"
	"trueabc.top/zinx/ziface"
	"trueabc.top/zinx/zinx_app_demo/mmo_game/game_server/core"
	"trueabc.top/zinx/zinx_app_demo/mmo_game/game_server/pb"
	"trueabc.top/zinx/znet"
)

type WorldChatApi struct {
	znet.BaseRouter
}

func (wc *WorldChatApi) Handle(request ziface.IRequest) {
	// 解析客户端的proto协议
	proto_msg := pb.Talk{}
	err := proto_msg.Unmarshal(request.GetData())
	if err != nil {
		fmt.Println("Task message unmarshall error")
		return
	}

	//
	pid, _ := request.GetConnection().GetProperty("pid")
	player := core.WManObj.GetPlayerByPid(pid.(int32))

	player.Talk(proto_msg.Content)

	// 根据pid得到对象

	// 将消息广播给其他在线的玩家

}
