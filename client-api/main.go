package main

import (
	"flag"
	"github.com/ktsymbal/seaports-catalogue/client-api/http/rest"
	"github.com/ktsymbal/seaports-catalogue/client-api/managing"
	"github.com/ktsymbal/seaports-catalogue/client-api/populating"
	"log"
	"os"
	"os/signal"

	"github.com/ktsymbal/seaports-catalogue/pkg/lifecycle"
)

var (
	repoAddress = getFromEnv("REPOSITORY", "repository:8080")
	dataFile    = getFromEnv("DATA_FILE", "/opt/api/data/ports.json")
	port        = getFromEnv("PORT", "8080")
)

func main() {
	flag.Parse()

	managingSvc, err := managing.NewService(repoAddress)
	if err != nil {
		log.Fatal(err)
	}

	populatingSvc := populating.NewService(dataFile, managingSvc)
	api := rest.NewClientAPI(managingSvc, port)

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

func getFromEnv(key, defaultValue string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return
}
