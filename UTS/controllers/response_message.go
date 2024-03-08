package controllers

import (
	"encoding/json"
	"net/http"

	m "pbp/UTS/models"
)

func SendSuccessResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response m.SuccessResponse
	response.Status = status
	response.Message = message
	json.NewEncoder(w).Encode(response)
}
func SendErrorResponse(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response m.ErrorResponse
	response.Status = status
	response.Message = message
	json.NewEncoder(w).Encode(response)
}
