package core

import (
	"github.com/floreabogdan/inverter-app/bms"
	"github.com/floreabogdan/inverter-app/grid"
	"github.com/floreabogdan/inverter-app/inverter"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("module", "core")

type Core struct {
	WebEngine *gin.Engine
	Bms       map[string]bms.Bms
	Grid      map[string]grid.Grid
	Inverter  map[string]inverter.Inverter
}

func New() *Core {
	return &Core{
		WebEngine: gin.Default(),
		Bms:       make(map[string]bms.Bms),
		Grid:      make(map[string]grid.Grid),
		Inverter:  make(map[string]inverter.Inverter),
	}
}

func (c *Core) Run() {
	inverterList := []map[string]interface{}{
		{
			"id":   "dummyInverter",
			"type": "dummy",

			"config":  "",
			"bms":     "dummyBms",
			"gridIn":  "dummyGridIn",
			"gridOut": "dummyGridOut",
		},
	}

	bmsList := []map[string]interface{}{
		{
			"id":   "dummyBms",
			"type": "dummy",

			"config": "",
		},
	}

	gridList := []map[string]interface{}{
		{
			"id":   "dummyGridIn",
			"type": "dummy",

			"config": "",
		},
		{
			"id":   "dummyGridOut",
			"type": "dummy",

			"config": "",
		},
	}

	// register bms
	c.Bms = bms.LoadBms(bmsList)

	// register grid
	c.Grid = grid.LoadGrid(gridList)

	// register inverters.
	c.Inverter = inverter.LoadInverters(inverterList, c.Bms, c.Grid)

	err := c.WebEngine.Run()
	if err != nil {
		log.Fatal(err)
	}
}
