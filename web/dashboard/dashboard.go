package dashboard

import (
	"github.com/floreabogdan/inverter-app/core"

	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("module", "dashboard")

type Dashboard struct {
	core *core.Core
}

func New(core *core.Core) *Dashboard {
	return &Dashboard{
		core: core,
	}
}

func (d *Dashboard) Run() {
	d.initRoutes()
}
