package main

import (
	"fmt"
	"trueabc.top/zinx/ziface"
	"trueabc.top/zinx/zinx_app_demo/mmo_game/game_server/apis"
	"trueabc.top/zinx/zinx_app_demo/mmo_game/game_server/core"
	"trueabc.top/zinx/znet"
)

func main() {
	s := znet.NewServer("MMO Game Server")

	s.SetOnConnStart(func(connection ziface.IConnection) {
		// 创建一个player
		player := core.NewPlayer(connection)

		// 发送MsgId：1消息和广播200消息
		player.SyncPid()

		player.BroadCastStartPosition()
		core.WManObj.AddPlayer(player)

		connection.SetProperty("pid", player.Pid)
		// 记录当前连接属于哪个玩家

		// 同步周边玩家, 广播当前玩家的位置信息
		player.SyncSurrounding()
		fmt.Println("------> player id ", player.Pid, " is arrived.")
	})

	s.SetOnConnStop(func(connection ziface.IConnection) {

	})
	// 注册路由业务
	s.AddRouter(2, &apis.WorldChatApi{}) // 聊天函数
	s.AddRouter(3, &apis.MoveApi{})      // 位置信息更新
	// 连接创建和销毁的钩子函数

	s.Serve()
}
