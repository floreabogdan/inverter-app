package dummy

import (
	"github.com/floreabogdan/inverter-app/bms"
	"github.com/floreabogdan/inverter-app/grid"
	"github.com/goccy/go-json"
	"log"
)

type Dummy struct {
	config  string
	bms     bms.Bms
	gridIn  grid.Grid
	gridOut grid.Grid
}

func New(config string, bms bms.Bms, gridIn grid.Grid, gridOut grid.Grid) *Dummy {
	return &Dummy{
		config:  config,
		bms:     bms,
		gridIn:  gridIn,
		gridOut: gridOut,
	}
}

func (d Dummy) SetConfig(config string) {
	print(config)
}

func (d Dummy) GetConfig() string {
	_json, err := json.Marshal(d.bms)
	if err != nil {
		log.Fatal(err)
	}

	return string(_json)
}
