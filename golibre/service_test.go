package golibre_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/equalsgibson/golibre/golibre"
)

func newTestServer(t *testing.T) *httptest.Server {
	t.Helper()

	notAuthenticatedError := golibre.ErrorResponse{
		Status: 2,
		Error: golibre.ErrorMessage{
			Message: "notAuthenticated",
		},
	}

	testServeMux := http.NewServeMux()
	testServeMux.Handle("/", http.NotFoundHandler())
	testServeMux.HandleFunc("/llu/auth/login", func(w http.ResponseWriter, r *http.Request) {
		// Validate that on login, we are not sending an Authorization header
		
		if r.Header.Get("Authorization") != "" {
			w.Write()
			w.WriteHeader(http.StatusOK)

			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusOK)

			return
		}

		target := &golibre.Authentication{}
		if err := json.Unmarshal(body, target); err != nil {
			w.WriteHeader(http.StatusOK)

			return
		}

		validAuth := 

	})

	return httptest.NewTLSServer(
		testServeMux)

}
