package response

type HealthCheckResponse struct {
	Health bool `json:"healthy"`
}