package rest

import (
	"context"
	"github.com/ktsymbal/seaports-catalogue/client-api/managing"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ktsymbal/seaports-catalogue/pkg/lifecycle"

	"github.com/gorilla/handlers"
)

const ShutdownTimeout = 5 * time.Second

type API struct {
	server *http.Server

	port string
	err  error
}

func NewClientAPI(manager managing.Service, port string) *API {
	handler := NewHandler(manager)
	server := &http.Server{Addr: ":" + port, Handler: handlers.LoggingHandler(os.Stdout, handler)}
	return &API{server: server, port: port}
}

// Run starts API HTTP server
func (api *API) Run() {
	log.Printf("API: Listening on %s...", api.port)
	if err := api.server.ListenAndServe(); err != http.ErrServerClosed {
		api.err = err
		lifecycle.KillTheApp()
	}
}

// Stop attempts to gracefully stop API server with timeout
func (api *API) Stop() (err error) {
	if api.err != nil {
		return api.err
	}
	if api.server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
		defer cancel()
		if err = api.server.Shutdown(ctx); err != nil {
			return err
		}
	}
	log.Print("API stopped.")
	return
}
