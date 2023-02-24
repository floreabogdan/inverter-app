package bms

import (
	"fmt"
	"github.com/floreabogdan/inverter-app/bms/dummy"
)

type Bms interface {
	SetConfig(string)
	GetConfig() string
}

func LoadBms(gridData map[string]map[string]string) map[string]Bms {
	bmsList := make(map[string]Bms)

	for key, bmsInfo := range gridData {
		bmsType := fmt.Sprint(bmsInfo["type"])
		config := fmt.Sprint(bmsInfo["config"])

		switch bmsType {
		case "dummy":
			bmsList[key] = dummy.New(config)
			break
		}
	}

	return bmsList
}
