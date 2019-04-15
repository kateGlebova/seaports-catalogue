package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/kateGlebova/seaports-catalogue/pkg/populating"

	"github.com/kateGlebova/seaports-catalogue/pkg/managing"

	"github.com/kateGlebova/seaports-catalogue/pkg/http/rest"

	"github.com/kateGlebova/seaports-catalogue/pkg/lifecycle"
)

func main() {
	managingSvc, err := managing.NewService(":9090")
	if err != nil {
		log.Fatal(err)
	}

	populatingSvc := populating.NewService("ports.json", managingSvc)
	api := rest.NewClientAPI(managingSvc, "8080")

	stopper := lifecycle.NewStopper(api, managingSvc.(lifecycle.Stoppable), populatingSvc.(lifecycle.Stoppable))
	runner := lifecycle.NewRunner(api, populatingSvc.(lifecycle.Runnable))

	signalChan := make(chan os.Signal, 1)
	exitChan := make(chan int)
	signal.Notify(signalChan, lifecycle.GracefulShutdownSignals...)
	go lifecycle.SignalHandle(signalChan, exitChan, stopper.Stop)
	go runner.Run()

	code := <-exitChan
	os.Exit(code)
}
