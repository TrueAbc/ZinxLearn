package core

import "sync"

/*
当前世界的管理模块
*/

type WorldManager struct {
	aoi *AOIManager

	Players map[int32]*Player
	pLock   sync.RWMutex
}

// 提供一个对外的世界管理模块
var WManObj *WorldManager

// 初始化方法
func init() {
	WManObj = &WorldManager{
		aoi:     NewAOIManager(AOI_MIN_X, AOI_MAX_X, AOI_MIN_Y, AOI_MAX_Y, AOI_CNTS_X, AOI_CNTS_Y),
		Players: make(map[int32]*Player),
	}
}

// 添加玩家
func (m *WorldManager) AddPlayer(player *Player) {
	m.pLock.Lock()
	defer m.pLock.Unlock()
	m.Players[player.Pid] = player

	// 将player添加到aoi中
	m.aoi.AddToGridByPos(int(player.Pid), player.X, player.Z)
}

func (m *WorldManager) RemovePlayer(pid int32) {
	m.pLock.RLock()
	p := m.Players[pid]
	m.pLock.RUnlock()

	m.aoi.RemoveFromGridbyPos(int(pid), p.X, p.Z)

	m.pLock.Lock()
	delete(m.Players, pid)
	m.pLock.Unlock()
}

func (m *WorldManager) GetPlayerByPid(pid int32) *Player {
	m.pLock.RLock()
	defer m.pLock.RUnlock()
	return m.Players[pid]
}

func (m *WorldManager) GetAllPlayers() []*Player {
	m.pLock.RLock()
	defer m.pLock.RUnlock()

	players := make([]*Player, 0)
	for _, g := range m.Players {
		players = append(players, g)
	}
	return players
}
