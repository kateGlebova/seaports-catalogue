package main

import (
	"os"
	"os/signal"

	"github.com/kateGlebova/seaports-catalogue/pkg/storage/mongo"

	"github.com/kateGlebova/seaports-catalogue/pkg/http/proto"

	"github.com/kateGlebova/seaports-catalogue/pkg/lifecycle"
)

var (
	url        = getFromEnv("MONGO_URL", "mongo:27017")
	db         = getFromEnv("MONGO_DB", "ports")
	collection = getFromEnv("MONGO_COLLECTION", "ports")
	port       = getFromEnv("PORT", "8080")
)

func main() {
	storage, err := mongo.NewRepository(url, db, collection)
	if err != nil {
		panic(err)
	}
	portDomainSvc := proto.NewPortDomainService(port, storage)

	stopper := lifecycle.NewStopper(portDomainSvc, storage)

	signalChan := make(chan os.Signal, 1)
	exitChan := make(chan int)
	signal.Notify(signalChan, lifecycle.GracefulShutdownSignals...)
	go lifecycle.SignalHandle(signalChan, exitChan, stopper.Stop)
	go portDomainSvc.Run()
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
