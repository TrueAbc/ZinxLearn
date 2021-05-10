package core

import "fmt"

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

// 打印格子信息
