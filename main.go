package main

import (
	"github.com/floreabogdan/inverter-app/core"
	"github.com/floreabogdan/inverter-app/web/api"
	"github.com/floreabogdan/inverter-app/web/dashboard"

	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

var log = logrus.WithField("module", "main")

func main() {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT)
	signal.Notify(stopChan, syscall.SIGTERM)

	core := core.New()
	go core.Run()

	api := api.New(core)
	go api.Run()

	dash := dashboard.New(core)
	go dash.Run()

	<-stopChan
	log.Info("Got stop signal. Finishing work.")
}
