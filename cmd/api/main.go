package main

import (
	"os"
	"os/signal"

	"github.com/kateGlebova/seaports-catalogue/pkg/http/rest"

	"github.com/kateGlebova/seaports-catalogue/internal/api"
	"github.com/kateGlebova/seaports-catalogue/pkg/shutdown"
)

func main() {
	parser := rest.MockParser{}
	retriever := rest.MockRetriever{}
	handler := rest.NewHandler(retriever)
	a := api.NewClientAPI(parser, retriever, handler, "8080")

	signalChan := make(chan os.Signal, 1)
	exitChan := make(chan int)
	signal.Notify(signalChan, shutdown.GracefulShutdownSignals...)
	go shutdown.SignalHandle(signalChan, exitChan, a.Stop)
	go a.Run()
	code := <-exitChan
	os.Exit(code)
}
