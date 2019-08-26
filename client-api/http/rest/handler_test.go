package rest

import (
	"bytes"
	"encoding/json"
	"github.com/kateGlebova/seaports-catalogue/client-api/managing"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/kateGlebova/seaports-catalogue/pkg/entities"
	"github.com/stretchr/testify/assert"
)

func TestGetPortsEmpty(t *testing.T) {
	manager := managing.MockService{Ports: []entities.Port{}}
	router := NewHandler(manager)

	req, _ := http.NewRequest(http.MethodGet, "/ports", nil)
	response := ExecuteRequest(router, req)

	assert.Equal(t, http.StatusOK, response.Code)
	expected, _ := json.Marshal([]entities.Port{})
	assert.Equal(t, string(expected)+"\n", response.Body.String())
}

func TestGetPorts(t *testing.T) {
	ports := entities.MockPorts(100)
	manager := managing.MockService{Ports: ports}
	router := NewHandler(manager)

	testCases := []struct {
		name  string
		query string
		code  int
		body  interface{}
	}{
		{name: "success", query: "?limit=10&offset=0", code: http.StatusOK, body: ports},
		{name: "invalid limit", query: "?limit=13f&offset=4", code: http.StatusBadRequest, body: Response{Code: http.StatusBadRequest, Message: `Bad Request: strconv.Atoi: parsing "13f": invalid syntax`}},
		{name: "invalid offset", query: "?limit=13&offset=4df", code: http.StatusBadRequest, body: Response{Code: http.StatusBadRequest, Message: `Bad Request: strconv.Atoi: parsing "4df": invalid syntax`}},
		{name: "no limit and offset", code: http.StatusOK, body: ports},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/ports"+tc.query, nil)
			response := ExecuteRequest(router, req)

			assert.Equal(t, tc.code, response.Code)
			expected, _ := json.Marshal(tc.body)
			assert.Equal(t, string(expected)+"\n", response.Body.String())
		})
	}
}

func TestGetPort(t *testing.T) {
	manager := managing.MockService{}
	router := NewHandler(manager)

	req, _ := http.NewRequest(http.MethodGet, "/ports/AEAJM", nil)
	response := ExecuteRequest(router, req)

	assert.Equal(t, http.StatusOK, response.Code)
	expected, _ := json.Marshal(entities.MockPort)
	assert.Equal(t, string(expected)+"\n", response.Body.String())
}

func TestCreatePort(t *testing.T) {
	manager := managing.MockService{}
	router := NewHandler(manager)

	testCases := []struct {
		name   string
		reader io.Reader
		code   int
		body   interface{}
	}{
		{name: "body reading error", reader: errReader{}, code: http.StatusBadRequest, body: Response{Code: http.StatusBadRequest, Message: "Bad Request: test error"}},
		{name: "unmarshalling error", reader: strings.NewReader(`{"id":"AEAJM"`), code: http.StatusBadRequest, body: Response{Code: http.StatusBadRequest, Message: "Bad Request: unexpected end of JSON input"}},
		{name: "success", reader: convertToReader(entities.MockPort), code: http.StatusCreated, body: entities.MockPort},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodPost, "/ports", tc.reader)
			response := ExecuteRequest(router, req)

			assert.Equal(t, tc.code, response.Code)
			expected, _ := json.Marshal(tc.body)
			assert.Equal(t, string(expected)+"\n", response.Body.String())
		})
	}
}

func TestUpdatePort(t *testing.T) {
	manager := managing.MockService{}
	router := NewHandler(manager)

	testCases := []struct {
		name   string
		id     string
		reader io.Reader
		code   int
		body   string
	}{
		{name: "body reading error", id: "AEAJM", reader: errReader{}, code: http.StatusBadRequest, body: "{\"code\":400,\"message\":\"Bad Request: test error\"}\n"},
		{name: "unmarshalling error", id: "AEAJM", reader: strings.NewReader(`{"id":"AEAJM"`), code: http.StatusBadRequest, body: "{\"code\":400,\"message\":\"Bad Request: unexpected end of JSON input\"}\n"},
		{name: "port ID mismatch", id: "AEAJMD", reader: convertToReader(entities.MockPort), code: http.StatusBadRequest, body: "{\"code\":400,\"message\":\"Bad Request: port ID in URL path does not match port ID in body\"}\n"},
		{name: "success", id: "AEAJM", reader: convertToReader(entities.MockPort), code: http.StatusNoContent, body: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodPut, "/ports/"+tc.id, tc.reader)
			response := ExecuteRequest(router, req)

			assert.Equal(t, tc.code, response.Code)
			assert.Equal(t, tc.body, response.Body.String())
		})
	}
}

func TestDeletePort(t *testing.T) {
	manager := managing.MockService{}
	router := NewHandler(manager)

	req, _ := http.NewRequest(http.MethodDelete, "/ports/AEAJM", nil)
	response := ExecuteRequest(router, req)

	assert.Equal(t, http.StatusNoContent, response.Code)
}

func TestServiceError(t *testing.T) {
	manager := managing.MockService{Err: testError{}}
	router := NewHandler(manager)
	testCases := []struct {
		name   string
		url    string
		method string
		body   interface{}
	}{
		{name: "list ports", url: "/ports", method: http.MethodGet},
		{name: "create port", url: "/ports", method: http.MethodPost, body: entities.MockPort},
		{name: "get port", url: "/ports/AEAJM", method: http.MethodGet},
		{name: "update port", url: "/ports/AEAJM", method: http.MethodPut, body: entities.MockPort},
		{name: "delete port", url: "/ports/AEAJM", method: http.MethodDelete},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var body io.Reader
			if tc.body != nil {
				body = convertToReader(tc.body)
			}
			req, _ := http.NewRequest(tc.method, tc.url, body)
			response := ExecuteRequest(router, req)

			assert.Equal(t, http.StatusInternalServerError, response.Code)
			expected, _ := json.Marshal(Response{Code: http.StatusInternalServerError, Message: "Internal Server Error: test error"})
			assert.Equal(t, string(expected)+"\n", response.Body.String())
		})
	}
}

func convertToReader(entity interface{}) io.Reader {
	payload, _ := json.Marshal(entity)
	return bytes.NewBuffer(payload)
}
