package api

import (
	"context"
	"log"
	"net/http"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/kateGlebova/seaports-catalogue/pkg/entities"
	"github.com/kateGlebova/seaports-catalogue/pkg/parser"
)

const ShutdownTimeout = 5 * time.Second

type ClientAPI struct {
	server     *http.Server
	parser     parser.Service
	repository entities.PortRepository

	port string
	err  error
}

func NewClientAPI(p parser.Service, r entities.PortRepository, port string) *ClientAPI {
	return &ClientAPI{parser: p, repository: r, port: port}
}

func (api *ClientAPI) InitialiseServer() {
	api.server = &http.Server{Addr: ":" + api.port, Handler: api.NewRouter()}
}

func (api *ClientAPI) NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ports", api.GetPorts)
	r.HandleFunc("/ports/{port}", api.GetPort)
	return r
}

// Run starts listening and serving incoming HTTP requests
func (api *ClientAPI) Run() {
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
			log.Fatalf("Stoping ClientAPI error: %v\n", err)
		} else {
			log.Print("ClientAPI stopped")
		}
	}
	return
}

// killTheApp sends SIGTERM to the parent application to quit
func killTheApp() {
	pid := syscall.Getpid()
	syscall.Kill(pid, syscall.SIGTERM)
}
