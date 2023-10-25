package controllers

import (
	"encoding/json"
	"net/http"

	response "github.com/backend-services/models/response"
)

// RespondWithError encapsulates a generic error response
func RespondWithError(w http.ResponseWriter, code int, message string, messageCode string) {
	data, _ := json.Marshal(response.ErrorResponse{Error: response.ErrorResponseError{Code: code, Message: message, MessageCode: messageCode}})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
