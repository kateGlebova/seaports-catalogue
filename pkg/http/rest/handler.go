package rest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kateGlebova/seaports-catalogue/pkg/entities"
	"github.com/kateGlebova/seaports-catalogue/pkg/managing"
)

func NewHandler(manager managing.Service) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/ports", getPorts(manager)).Methods(http.MethodGet)
	r.HandleFunc("/ports", getPorts(manager)).Queries("limit", "{limit}", "offset", "{offset}").Methods(http.MethodGet)
	r.HandleFunc("/ports", createPort(manager)).Methods(http.MethodPost)
	r.HandleFunc("/ports/{port}", getPort(manager)).Methods(http.MethodGet)
	r.HandleFunc("/ports/{port}", updatePort(manager)).Methods(http.MethodPut, http.MethodPatch)
	r.HandleFunc("/ports/{port}", deletePort(manager)).Methods(http.MethodDelete)
	return r
}

func getPorts(manager managing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, offset, err := parseLimitAndOffset(r)
		if err != nil {
			BadRequest(w, err)
			return
		}

		ports, err := manager.ListPorts(uint(limit), uint(offset))
		if err != nil {
			Error(w, err)
			return
		}
		SuccessWithEntity(w, ports)
	}
}

func getPort(manager managing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		portID := params["port"]

		port, err := manager.GetPort(portID)
		if err != nil {
			Error(w, err)
			return
		}
		SuccessWithEntity(w, port)
	}
}

func createPort(manager managing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			BadRequest(w, err)
			return
		}
		defer r.Body.Close()

		var port entities.Port
		err = json.Unmarshal(body, &port)
		if err != nil {
			BadRequest(w, err)
			return
		}

		err = manager.CreatePort(port)
		if err != nil {
			Error(w, err)
			return
		}

		Created(w, port)
	}
}

func updatePort(manager managing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		portID := params["port"]

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			BadRequest(w, err)
			return
		}
		defer r.Body.Close()

		var port entities.Port
		err = json.Unmarshal(body, &port)
		if err != nil {
			BadRequest(w, err)
			return
		}

		if port.ID != "" && portID != port.ID {
			BadRequest(w, errors.New("port ID in URL path does not match port ID in body"))
			return
		}

		err = manager.UpdatePort(portID)
		if err != nil {
			Error(w, err)
			return
		}
		SuccessWithEntity(w, port)
	}
}

func deletePort(manager managing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		portID := params["port"]

		err := manager.DeletePort(portID)
		if err != nil {
			Error(w, err)
			return
		}

		NoContent(w)
	}
}

func parseLimitAndOffset(r *http.Request) (limit, offset int, err error) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	if limitStr == "" {
		limit = 0
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return
		}
	}

	if offsetStr == "" {
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetStr)
	}

	return
}
