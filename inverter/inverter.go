package inverter

import (
	"fmt"
	"github.com/floreabogdan/inverter-app/bms"
	"github.com/floreabogdan/inverter-app/grid"
	"github.com/floreabogdan/inverter-app/inverter/dummy"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("module", "core")

type Inverter interface {
	SetConfig(string)
	GetConfig() string
}

func LoadInverters(invertersData []map[string]interface{}, bmsList map[string]bms.Bms, gridList map[string]grid.Grid) map[string]Inverter {
	invertersList := make(map[string]Inverter)

	for _, inv := range invertersData {
		id := fmt.Sprint(inv["id"])
		inverterType := fmt.Sprint(inv["type"])
		config := fmt.Sprint(inv["config"])

		bmsId := fmt.Sprint(inv["bms"])
		gridInId := fmt.Sprint(inv["gridIn"])
		gridOutId := fmt.Sprint(inv["gridOut"])

		_bms, ok := bmsList[bmsId]
		if !ok {
			log.Error("Could not find BMS")
		}

		_gridIn, ok := gridList[gridInId]
		if !ok {
			log.Error("Could not find GRID IN")
		}

		_gridOut, ok := gridList[gridOutId]
		if !ok {
			log.Error("Could not find GRID OUT")
		}

		switch inverterType {
		case "dummy":
			invertersList[id] = dummy.New(config, _bms, _gridIn, _gridOut)
			break
		}
	}

	return invertersList
}
