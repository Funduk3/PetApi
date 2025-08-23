package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SendResponse(w http.ResponseWriter, statusCode int, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func SendSuccessResponse(w http.ResponseWriter, data interface{}, message string) {
	SendResponse(w, http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func SendCreatedResponse(w http.ResponseWriter, data interface{}, message string) {
	SendResponse(w, http.StatusCreated, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func SendErrorResponse(w http.ResponseWriter, statusCode int, errorMsg string) {
	SendResponse(w, statusCode, Response{
		Success: false,
		Error:   errorMsg,
	})
}
