package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
	"github.com/kateGlebova/seaports-catalogue/pkg/parsing"
	"github.com/kateGlebova/seaports-catalogue/pkg/retrieving"
)

const ShutdownTimeout = 5 * time.Second

type ClientAPI struct {
	server    *http.Server
	parser    parsing.Service
	retriever retrieving.Service

	port string
	err  error
}

func NewClientAPI(p parsing.Service, r retrieving.Service, handler http.Handler, port string) *ClientAPI {
	server := &http.Server{Addr: ":" + port, Handler: handlers.LoggingHandler(os.Stdout, handler)}
	return &ClientAPI{parser: p, retriever: r, server: server, port: port}
}

// Run starts listening and serving incoming HTTP requests
func (api *ClientAPI) Run() {
	log.Printf("Listening on %s...", api.port)
	if err := api.server.ListenAndServe(); err != http.ErrServerClosed {
		api.err = err
		killTheApp()
	}
}

//Stop attempts to gracefully stop API server with timeout
func (api *ClientAPI) Stop() (err error) {
	if api.err != nil {
		return api.err
	}
	if api.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
		defer cancel()
		if err = api.server.Shutdown(ctx); err != nil {
			return err
		}
		log.Print("ClientAPI stopped")
	}
	return
}

// killTheApp sends SIGTERM to the parent application to quit
func killTheApp() {
	pid := syscall.Getpid()
	syscall.Kill(pid, syscall.SIGTERM)
}
