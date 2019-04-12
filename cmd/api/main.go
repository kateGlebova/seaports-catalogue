package main

import (
	"os"
	"os/signal"

	"github.com/kateGlebova/seaports-catalogue/internal/api"
	"github.com/kateGlebova/seaports-catalogue/pkg/shutdown"
)

func main() {
	parser := api.MockParser{}
	repo := api.MockRepo{}
	a := api.NewClientAPI(parser, repo, "8080")
	a.InitialiseServer()
	//a.Run()

	signalChan := make(chan os.Signal, 1)
	exitChan := make(chan int)
	signal.Notify(signalChan, shutdown.GracefulShutdownSignals...)
	go shutdown.SignalHandle(signalChan, exitChan, a.Stop)
	go a.Run()
	code := <-exitChan
	os.Exit(code)
}
