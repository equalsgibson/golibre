package golibre

import (
	"fmt"
	"net/http"
)

type APIError struct {
	RawResponse *http.Response `json:"rawResponse"`
	Status      StatusCode     `json:"status"`
	Detail      ErrorMessage   `json:"error"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func (e *APIError) Error() string {
	if e.Detail.Message == "" {
		return fmt.Sprintf("generic LibreView API error, network status code: %d", e.RawResponse.StatusCode)
	}

	return e.Detail.Message
}
