package main

import (
	"os"
	"os/signal"

	"github.com/kateGlebova/seaports-catalogue/pkg/storage/mongo"

	"github.com/kateGlebova/seaports-catalogue/pkg/http/proto"

	"github.com/kateGlebova/seaports-catalogue/pkg/lifecycle"
)

func main() {
	storage := mongo.NewRepository("localhost:27017", "ports", "ports")
	portDomainSvc := proto.NewPortDomainService("9090", storage)

	runner := lifecycle.NewRunner(portDomainSvc, storage)

	signalChan := make(chan os.Signal, 1)
	exitChan := make(chan int)
	signal.Notify(signalChan, lifecycle.GracefulShutdownSignals...)
	go lifecycle.SignalHandle(signalChan, exitChan, runner.Stop)
	go runner.Run()
	code := <-exitChan
	os.Exit(code)
}
