package apis

import (
	"fmt"
	"trueabc.top/zinx/ziface"
	"trueabc.top/zinx/zinx_app_demo/mmo_game/game_server/core"
	"trueabc.top/zinx/zinx_app_demo/mmo_game/game_server/pb"
	"trueabc.top/zinx/znet"
)

type MoveApi struct {
	znet.BaseRouter
}

func (a *MoveApi) Handler(request ziface.IRequest) {
	// 解析客户端的协议
	proto_msg := &pb.Position{}
	err := proto_msg.Unmarshal(request.GetData())
	if err != nil {
		fmt.Println("Move: Position Unmarshal error ", err)
		return
	}
	// 得到当前发送位置的玩家信息
	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		fmt.Println("GetProperty pid error: ", err)
		return
	}

	fmt.Printf("User pid = %d, move(%f, %f, %f, %f)\n", pid, proto_msg.X, proto_msg.Y, proto_msg.Z, proto_msg.V)

	// 给其他玩家进行当前玩家的位置信息广播
	player := core.WManObj.GetPlayerByPid(pid.(int32))
	player.UpdatePos(proto_msg.X, proto_msg.Y, proto_msg.Z, proto_msg.V)

}
