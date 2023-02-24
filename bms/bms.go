package bms

import (
	"fmt"
	"github.com/floreabogdan/inverter-app/bms/dummy"
)

type Bms interface {
	SetConfig(string)
	GetConfig() string
}

func LoadBms(gridData []map[string]interface{}) map[string]Bms {
	bmsList := make(map[string]Bms)

	for _, inv := range gridData {
		id := fmt.Sprint(inv["id"])
		bmsType := fmt.Sprint(inv["type"])
		config := fmt.Sprint(inv["config"])

		switch bmsType {
		case "dummy":
			bmsList[id] = dummy.New(config)
			break
		}
	}

	return bmsList
}
