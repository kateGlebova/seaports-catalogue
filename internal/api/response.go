package api

// Response represents response body
type Response struct {
	// HTTP status code
	Code int `json:"code"`
	// Response message
	Message string `json:"message"`
}
