package main

import (
	"os"
	"os/signal"

	"github.com/kateGlebova/seaports-catalogue/pkg/storage/mongo"

	"github.com/kateGlebova/seaports-catalogue/pkg/http/proto"

	"github.com/kateGlebova/seaports-catalogue/pkg/lifecycle"
)

func main() {
	storage, err := mongo.NewRepository("localhost:27017", "ports", "ports")
	if err != nil {
		panic(err)
	}
	portDomainSvc := proto.NewPortDomainService("9090", storage)

	stopper := lifecycle.NewStopper(portDomainSvc, storage)

	signalChan := make(chan os.Signal, 1)
	exitChan := make(chan int)
	signal.Notify(signalChan, lifecycle.GracefulShutdownSignals...)
	go lifecycle.SignalHandle(signalChan, exitChan, stopper.Stop)
	go portDomainSvc.Run()
	code := <-exitChan
	os.Exit(code)
}
