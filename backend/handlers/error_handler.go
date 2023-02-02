package handlers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func respondWithError(w http.ResponseWriter, message string, code int) {
	ErrorResponse := ErrorResponse{Message: message, Code: code}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse)
}
