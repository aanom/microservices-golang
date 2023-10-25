package response

// ErrorResponse -
type ErrorResponse struct {
	Error ErrorResponseError `json:"error"`
}

// ErrorResponseError -
type ErrorResponseError struct {
	Code        int    `json:"code"`
	MessageCode string `json:"messageCode"`
	Message     string `json:"message"`
}
