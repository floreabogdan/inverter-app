package grid

import (
	"fmt"
	"github.com/floreabogdan/inverter-app/grid/dummy"
)

type Grid interface {
	SetConfig(string)
	GetConfig() string
}

func LoadGrid(gridData []map[string]interface{}) map[string]Grid {
	gridList := make(map[string]Grid)

	for _, inv := range gridData {
		id := fmt.Sprint(inv["id"])
		gridType := fmt.Sprint(inv["type"])
		config := fmt.Sprint(inv["config"])

		switch gridType {
		case "dummy":
			gridList[id] = dummy.New(config)
			break
		}
	}

	return gridList
}
