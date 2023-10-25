package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/backend-services/models/response"
)

// HealthCheck -
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	var result = response.HealthCheckResponse{}
	result.Health = true
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(result)
	w.Write(data)
	return
}
