package response

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func Success(w http.ResponseWriter, status int, message string) {
	JSON(w, status, SuccessResponse{
		Success: true,
		Message: message,
	})
}

func Error(w http.ResponseWriter, status int, message string, errors map[string]string) {
	JSON(w, status, ErrorResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}