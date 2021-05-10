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
