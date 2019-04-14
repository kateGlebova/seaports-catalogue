package main

import (
	"os"
	"os/signal"

	"github.com/kateGlebova/seaports-catalogue/pkg/populating"

	"github.com/kateGlebova/seaports-catalogue/pkg/managing"

	"github.com/kateGlebova/seaports-catalogue/pkg/http/rest"

	"github.com/kateGlebova/seaports-catalogue/pkg/lifecycle"
)

func main() {
	managingSvc := managing.NewService(":9090")
	populatingSvc := populating.NewService("ports.json", managingSvc)
	a := rest.NewClientAPI(managingSvc, "8080")

	runner := lifecycle.NewRunner(a, managingSvc.(lifecycle.Runnable), populatingSvc.(lifecycle.Runnable))

	signalChan := make(chan os.Signal, 1)
	exitChan := make(chan int)
	signal.Notify(signalChan, lifecycle.GracefulShutdownSignals...)
	go lifecycle.SignalHandle(signalChan, exitChan, runner.Stop)
	go runner.Run()
	code := <-exitChan
	os.Exit(code)
}
