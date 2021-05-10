package core

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"sync"
	"trueabc.top/zinx/ziface"
	"trueabc.top/zinx/zinx_app_demo/mmo_game/game_server/pb"
)

type Player struct {
	Pid  int32
	Conn ziface.IConnection
	X    float32
	Y    float32
	Z    float32
	V    float32 // 角度坐标
}

var PidGen int32 = 1  // 用于生产玩家ID的计数器
var IdLock sync.Mutex // Mutex

func NewPlayer(conn ziface.IConnection) *Player {
	// 生成玩家Id
	IdLock.Lock()
	id := PidGen
	PidGen++
	IdLock.Unlock()

	//
	return &Player{
		Pid:  id,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)),
		Y:    0,
		Z:    float32(120 + rand.Intn(20)),
		V:    0, // 暂时没有用处
	}
}

/*
提供一个发送客户端消息的方法
将pb的protobuf的数据序列化之后, 再调用zinx
*/
func (p *Player) SendMsg(msgId uint32, data proto.Message) {

	// 将proto转化为二进制
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("marsahl msg err: ", err)
		return
	}

	// 将二进制通过zinx进行发送
	if p.Conn == nil {
		fmt.Println("connection in player is nil")
		return
	}

	if err := p.Conn.SendMsg(msgId, msg); err != nil {
		fmt.Println("send err ", err)
		return
	}

}

// 告知客户端玩家Id, 同步已经生成的玩家Id
func (p *Player) SyncPid() {
	data := &pb.SyncPid{
		Pid: p.Pid,
	}
	p.SendMsg(1, data)
}

// 广播玩家的出生地点
func (p *Player) BroadCastStartPosition() {
	data := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2, // 广播位置
		Data: &pb.BroadCast_P{
			// Position
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	p.SendMsg(1, data)
}
