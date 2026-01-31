package utils

import (
	"encoding/json"
	"net/http"
)

type DataResponse struct {
	Data interface{} `json:"data"`
}

type MessageResponse struct {
	Message interface{} `json:"message"`
}

func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(DataResponse{
		Data: data,
	})
}

func JSONMessageResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(MessageResponse{
		Message: message,
	})
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(MessageResponse{
		Message: message,
	})
}
