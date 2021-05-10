package main

import (
	"fmt"
	"trueabc.top/zinx/ziface"
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
		fmt.Println("------> player id ", player.Pid, " is arrived.")
	})

	s.SetOnConnStop(func(connection ziface.IConnection) {

	})
	// 注册路由业务

	// 连接创建和销毁的钩子函数

	s.Serve()
}
