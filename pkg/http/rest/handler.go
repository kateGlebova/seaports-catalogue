package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kateGlebova/seaports-catalogue/pkg/retrieving"
)

func NewHandler(retriever retrieving.Service) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/ports", getPorts(retriever)).Methods(http.MethodGet)
	r.HandleFunc("/ports/{port}", getPort(retriever)).Methods(http.MethodGet)
	return r
}

func getPorts(retriever retrieving.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ports := retriever.RetrieveAllPorts()
		err := json.NewEncoder(w).Encode(ports)
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

func getPort(retriever retrieving.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		port := retriever.RetrievePort("AEAJM")
		err := json.NewEncoder(w).Encode(port)
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
