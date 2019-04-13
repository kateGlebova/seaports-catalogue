package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kateGlebova/seaports-catalogue/pkg/managing"
)

func NewHandler(manager managing.Service) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/ports", getPorts(manager)).Methods(http.MethodGet)
	r.HandleFunc("/ports/{port}", getPort(manager)).Methods(http.MethodGet)
	return r
}

func getPorts(manager managing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ports, err := manager.ListPorts(10, 0)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			msg := fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError), err)
			json.NewEncoder(w).Encode(Response{
				Code:    http.StatusInternalServerError,
				Message: msg,
			})
			return
		}
		err = json.NewEncoder(w).Encode(ports)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			msg := fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError), err)
			json.NewEncoder(w).Encode(Response{
				Code:    http.StatusInternalServerError,
				Message: msg,
			})
			return
		}
	}
}

func getPort(manager managing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		port, err := manager.GetPort("AEAJM")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			msg := fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError), err)
			json.NewEncoder(w).Encode(Response{
				Code:    http.StatusInternalServerError,
				Message: msg,
			})
			return
		}
		err = json.NewEncoder(w).Encode(port)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			msg := fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError), err)
			json.NewEncoder(w).Encode(Response{
				Code:    http.StatusInternalServerError,
				Message: msg,
			})
			return
		}
	}
}
