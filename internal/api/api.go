package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kateGlebova/seaports-catalogue/pkg/lifecycle"

	"github.com/gorilla/handlers"
)

const ShutdownTimeout = 5 * time.Second

type ClientAPI struct {
	server *http.Server

	port string
	err  error
}

func NewClientAPI(handler http.Handler, port string) *ClientAPI {
	server := &http.Server{Addr: ":" + port, Handler: handlers.LoggingHandler(os.Stdout, handler)}
	return &ClientAPI{server: server, port: port}
}

// Run starts ClientAPI HTTP server
func (api *ClientAPI) Run() {
	log.Printf("Listening on %s...", api.port)
	if err := api.server.ListenAndServe(); err != http.ErrServerClosed {
		api.err = err
		lifecycle.KillTheApp()
	}
}

// Stop attempts to gracefully stop API server with timeout
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
