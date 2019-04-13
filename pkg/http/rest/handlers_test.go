// +build unit

package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kateGlebova/seaports-catalogue/pkg/entities"

	"github.com/gorilla/mux"
)

var clientAPI *ClientAPI

func TestMain(m *testing.M) {
	retriever := MockRetriever{}
	router := NewHandler(retriever)

	m.Run()
}

func TestClientAPI_GetPorts(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/ports", nil)
	response := ExecuteRequest(router, req)

	CheckResponseCode(t, http.StatusOK, response.Code)

	var ports []entities.Port
	json.Unmarshal(response.Body.Bytes(), &nodes)
	if len(nodes) != 0 {
		t.Errorf("Expected an empty array. Got %v", nodes)
	}
}

func ExecuteRequest(r *mux.Router, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	return rr
}

func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
