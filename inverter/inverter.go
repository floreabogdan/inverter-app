package inverter

import (
	"fmt"
	"github.com/floreabogdan/inverter-app/inverter/dummy"
	"github.com/floreabogdan/inverter-app/inverter/models"
)

type Inverter interface {
	SetConfig(models.Config)
	GetConfig() models.Config

	GetData() models.Data
	RunCommand(string) string
}

func LoadInverters(invertersData map[string]map[string]string) map[string]Inverter {
	invertersList := make(map[string]Inverter)

	for key, invInfo := range invertersData {
		inverterType := fmt.Sprint(invInfo["type"])
		config := fmt.Sprint(invInfo["config"])

		switch inverterType {
		case "dummy":
			invertersList[key] = dummy.New(models.Config{Connection: config})
			break
		}
	}

	return invertersList
}
