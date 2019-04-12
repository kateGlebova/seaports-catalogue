package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (api *ClientAPI) GetPorts(w http.ResponseWriter, r *http.Request) {
	ports := api.repository.GetAll()
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

func (api *ClientAPI) GetPort(w http.ResponseWriter, r *http.Request) {
	port := api.repository.Get("AEAJM")
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
