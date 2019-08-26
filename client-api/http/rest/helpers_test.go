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

type testError struct{}

func (err testError) Error() string {
	return "test error"
}

type errReader struct{}

func (r errReader) Read([]byte) (int, error) {
	return 0, testError{}
}
