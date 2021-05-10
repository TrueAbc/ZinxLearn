package core

import (
	"fmt"
)

/*
AOI区域管理模块
*/

type AOIManager struct {
	MinX int
	MinY int
	MaxX int
	MaxY int

	// x和y方向的格子数量
	CntsX int
	CntsY int

	// 格子id, 格子对象
	grids map[int]*Grid
}

func (m *AOIManager) String() string {
	s := fmt.Sprintf("AOIManager: \n MinX: %d, MaxX: %d, cntsX:%d \n"+
		"MinY:%d, MaxY:%d, cntsY: %d\n", m.MinX, m.MaxX, m.CntsX, m.MinY, m.MaxY, m.CntsY)

	for _, v := range m.grids {
		s += fmt.Sprintln(v)
	}

	return s
}

func NewAOIManager(minX, maxX, minY, maxY, cntsX, cntsY int) *AOIManager {

	a := &AOIManager{
		MinX: minX, MaxX: maxX, CntsX: cntsX,
		MinY: minY, MaxY: maxY, CntsY: cntsY,
		grids: make(map[int]*Grid),
	}
	for y := 0; y < cntsY; y++ {
		for x := 0; x < cntsX; x++ {
			// 格子编号: id = idy * cntX + idx
			id := y*cntsX + x
			a.grids[id] = NewGrid(id,
				a.MinX+x*a.gridWidth(),
				a.MinX+(x+1)*a.gridWidth(),
				a.MinY+(y)*a.gridLength(),
				a.MinY+(y+1)*a.gridLength())

		}
	}

	return a
}

// 得到每个格子在x和y轴方向的高度和宽度
func (m *AOIManager) gridWidth() int {
	return (m.MaxX - m.MinX) / m.CntsX
}

func (m *AOIManager) gridLength() int {
	return (m.MaxY - m.MinY) / m.CntsY
}

// 根据格子的gid得到周围的格子ID集合
func (m *AOIManager) GetSurroundGridsById(gID int) (grids []*Grid) {
	//判断当前的gid是否存在
	if _, ok := m.grids[gID]; !ok {
		return
	}

	// 将自己加入到切片中
	grids = append(grids, m.grids[gID])

	idx := gID % m.CntsX
	if idx > 0 {
		grids = append(grids, m.grids[gID-1])
	}

	if idx < m.CntsX-1 {
		grids = append(grids, m.grids[gID+1])
	}
	// 将x方向的点加入, 之后处理y轴的点

	// for 的range的特点
	for _, v := range grids {

		idy := v.GID / m.CntsX
		// 需要判断id是否是上下左右的边界格子
		if idy > 0 {
			grids = append(grids, m.grids[v.GID-m.CntsX])
		}
		if idy < m.CntsY-1 {
			grids = append(grids, m.grids[v.GID+m.CntsX])
		}
	}

	return grids
}

func (m *AOIManager) GetGIdByPos(x, y float32) int {
	ix := (int(x) - m.MinX) / m.gridWidth()
	iy := (int(y) - m.MinY) / m.gridLength()

	return iy*m.CntsX + ix
}

// 通过横纵坐标得到周边九宫格内全部的PlayerIDS
func (m *AOIManager) GetPIdsByPos(x, y float32) (playIDs []int) {
	// 得到玩家所在的格子id
	id := m.GetGIdByPos(x, y)
	// 通过gid得到周围的九宫格
	grids := m.GetSurroundGridsById(id)
	// 将九宫格里的全部player的id累加到playerIDs
	for _, v := range grids {
		playIDs = append(playIDs, v.GetPlayerIDs()...)
	}

	return playIDs
}
