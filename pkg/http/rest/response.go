package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"testing"
)

// Response represents response body
type Response struct {
	// HTTP status code
	Code int `json:"code"`
	// Response message
	Message string `json:"message"`
}

// Error returns error response with 500 code and error message
func Error(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	msg := fmt.Sprintf("%s: %s", http.StatusText(http.StatusInternalServerError), err)
	log.Printf("[%d] %s", http.StatusInternalServerError, msg)
	printStack()
	json.NewEncoder(w).Encode(Response{
		Code:    http.StatusInternalServerError,
		Message: msg,
	})
}

// NotFound returns error response with 404 code and "Not found" message
func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Response{
		Code:    http.StatusNotFound,
		Message: http.StatusText(http.StatusNotFound),
	})
}

// BadRequest returns error response with 400 code and "BadRequest: %s" error message
func BadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	msg := fmt.Sprintf("%s: %s", http.StatusText(http.StatusBadRequest), err)
	log.Printf("[%d] %s", http.StatusBadRequest, msg)
	json.NewEncoder(w).Encode(Response{
		Code:    http.StatusBadRequest,
		Message: msg,
	})
}

// Created returns response with 201 code and created entity
func Created(w http.ResponseWriter, entity interface{}) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entity)
}

// Created returns response with 204 code
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// SuccessWithEntity returns response with 200 code and JSON entity
func SuccessWithEntity(w http.ResponseWriter, entity interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entity)
}

func printStack() {
	// skip printing stack in tests in verbose mode
	if !testing.Verbose() {
		debug.PrintStack()
	}
}
