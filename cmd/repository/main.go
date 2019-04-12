package main

import (
	"os"
	"os/signal"

	"github.com/kateGlebova/seaports-catalogue/pkg/shutdown"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	exitChan := make(chan int)
	signal.Notify(signalChan, shutdown.GracefulShutdownSignals...)
	go shutdown.SignalHandle(signalChan, exitChan, func() error { return nil })
	//go start app
	code := <-exitChan
	os.Exit(code)
}
