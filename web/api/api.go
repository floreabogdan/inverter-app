package api

import (
	"github.com/floreabogdan/inverter-app/core"
	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("module", "api")

type Api struct {
	core *core.Core
}

func New(core *core.Core) *Api {
	return &Api{
		core: core,
	}
}

func (d *Api) Run() {
	d.initRoutes()
}
