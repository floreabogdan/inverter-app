package grid

import (
	"fmt"
	"github.com/floreabogdan/inverter-app/grid/dummy"
)

type Grid interface {
	SetConfig(string)
	GetConfig() string
}

func LoadGrid(gridData map[string]map[string]string) map[string]Grid {
	gridList := make(map[string]Grid)

	for key, gridInfo := range gridData {
		gridType := fmt.Sprint(gridInfo["type"])
		config := fmt.Sprint(gridInfo["config"])

		switch gridType {
		case "dummy":
			gridList[key] = dummy.New(config)
			break
		}
	}

	return gridList
}
