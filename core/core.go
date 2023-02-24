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

	Devices map[string]map[string]map[string]string
}

func New() *Core {
	return &Core{
		WebEngine: gin.Default(),
		Bms:       make(map[string]bms.Bms),
		Grid:      make(map[string]grid.Grid),
		Inverter:  make(map[string]inverter.Inverter),

		Devices: map[string]map[string]map[string]string{
			"inverter": {
				"dummyInverter": {
					"type":   "dummy",
					"config": "",

					"bms":     "dummyBms",
					"gridIn":  "dummyGridIn",
					"gridOut": "dummyGridOut",

					"mqttHost": "",
					"mqttPort": "",
					"mqttUser": "",
					"mqttPass": "",

					"influxHost":     "",
					"influxPort":     "",
					"influxUser":     "",
					"influxPass":     "",
					"influxDatabase": "",
				},
			},
			"bms": {
				"dummyBms": {
					"type":   "dummy",
					"config": "",

					"mqttHost": "",
					"mqttPort": "",
					"mqttUser": "",
					"mqttPass": "",

					"influxHost":     "",
					"influxPort":     "",
					"influxUser":     "",
					"influxPass":     "",
					"influxDatabase": "",
				},
			},
			"grid": {
				"dummyGridIn": {
					"type":   "dummy",
					"config": "",

					"mqttHost": "",
					"mqttPort": "",
					"mqttUser": "",
					"mqttPass": "",

					"influxHost":     "",
					"influxPort":     "",
					"influxUser":     "",
					"influxPass":     "",
					"influxDatabase": "",
				},
				"dummyGridOut": {
					"type":   "dummy",
					"config": "",

					"mqttHost": "",
					"mqttPort": "",
					"mqttUser": "",
					"mqttPass": "",

					"influxHost":     "",
					"influxPort":     "",
					"influxUser":     "",
					"influxPass":     "",
					"influxDatabase": "",
				},
			},
		},
	}
}

func (c *Core) Run() {
	//@todo setup local db

	//@todo setup mqtt

	//@todo setup influxdb

	//@todo load devices list

	// register bms.
	c.Bms = bms.LoadBms(c.Devices["bms"])

	// register grid.
	c.Grid = grid.LoadGrid(c.Devices["grid"])

	// register inverters.
	c.Inverter = inverter.LoadInverters(c.Devices["inverter"])

	//@todo start sync services

	// start web engine.
	err := c.WebEngine.Run()
	if err != nil {
		log.Fatal(err)
	}
}
