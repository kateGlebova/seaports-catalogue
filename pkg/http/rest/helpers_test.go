package rest

import (
	"net/http"
	"net/http/httptest"
)

func ExecuteRequest(r http.Handler, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	return rr
}
