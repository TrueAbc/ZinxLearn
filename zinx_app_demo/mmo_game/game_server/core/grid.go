package core

import (
	"fmt"
	"sync"
)

/*
AOI地图中的格子类型
*/

type Grid struct {
	GID int
	// 上下左右边界
	MinX int
	MinY int
	MaxX int
	MaxY int

	// 格子内的物体集合
	playerIDs map[int]bool
	// 集合的锁
	pIDLock sync.RWMutex
}

func NewGrid(gID, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		GID:       gID,
		MinX:      minX,
		MinY:      minY,
		MaxX:      maxX,
		MaxY:      maxY,
		playerIDs: make(map[int]bool),
	}
}

func (g *Grid) Add(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()

	g.playerIDs[playerID] = true
}

func (g *Grid) Remove(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()

	delete(g.playerIDs, playerID)
}

func (g *Grid) GetPlayerIDs() (playerIDs []int) {
	g.pIDLock.RLock()
	defer g.pIDLock.RUnlock()
	for k, _ := range g.playerIDs {
		playerIDs = append(playerIDs, k)
	}
	return playerIDs
}

func (g *Grid) String() string {
	return fmt.Sprintf("{Gid: %d, (minx, miny, maxx, maxy): (%d, %d, %d, %d), players: %v}",
		g.GID,
		g.MinX, g.MinY, g.MaxX, g.MaxY,
		g.playerIDs)
}
