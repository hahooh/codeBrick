package controllers

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Message string
}

// SetHeader set response header
func SetHeader(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
}

// JSONErrorResponse json response error
func JSONErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	SetHeader(w, statusCode)
	json.NewEncoder(w).Encode(errorResponse{Message: message})
}

// JSONResponse json response
func JSONResponse(w http.ResponseWriter, r *http.Request, responseData interface{}, statusCode int) {
	SetHeader(w, statusCode)
	json.NewEncoder(w).Encode(responseData)
}
