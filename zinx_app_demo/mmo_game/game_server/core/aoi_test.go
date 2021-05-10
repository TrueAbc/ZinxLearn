package core

import (
	"fmt"
	"testing"
)

func TestNewAOIManager(t *testing.T) {
	// 初始化AOIManager
	aoiModel := NewAOIManager(0, 100, 20, 200, 4, 1)
	fmt.Println(aoiModel)

	// 打印AOIManager
}

func TestAOIManager_GetSurroundGridsById(t *testing.T) {
	aoiModel := NewAOIManager(0, 250, 0, 250, 5, 5)
	fmt.Print(aoiModel.GetSurroundGridsById(0))
	for k, _ := range aoiModel.grids {
		//得到当前id的周边九宫格
		grids := aoiModel.GetSurroundGridsById(k)
		for _, item := range grids {
			fmt.Print(item.GID, " ")
		}
		fmt.Println()
	}
}
