package golibre_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/equalsgibson/golibre/golibre"
)

const validPassword string = "VALID_PASSWORD"
const validEmail string = "EMAIL"
const validJWTToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4LTEyMzQtMTIzNC0xMjM0LTQ2MDRhNjdiN2FiNCIsImZpcnN0TmFtZSI6IlNvbWUiLCJsYXN0TmFtZSI6Ik9uZSIsImNvdW50cnkiOiJHQiIsInJlZ2lvbiI6ImV1MiIsInJvbGUiOiJwYXRpZW50IiwidW5pdHMiOjAsInByYWN0aWNlcyI6W10sImMiOjEsInMiOiJsbHUuYW5kcm9pZCIsImV4cCI6MTc0MjI5NDIwN30.ilRwCINRf6nQViQ9c0BLZD9x21qsiBx43EzMk1POTuk"

func newTestServer(t *testing.T) (*httptest.Server, error) {
	t.Helper()

	testServeMux := http.NewServeMux()

	testServeMux.HandleFunc("/", http.NotFound)

	// Authentication endpoint
	testServeMux.HandleFunc("/llu/auth/login", func(w http.ResponseWriter, r *http.Request) {
		// Validate that on login, we are not sending an Authorization header
		if r.Header.Get("Authorization") != "" {
			t.Logf("Authorization header was present, with value: %s", r.Header.Get("Authorization"))
			notAuthenticated(w)

			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Logf("Error reading Request Body: %s", err.Error())
			notAuthenticated(w)

			return
		}

		target := &golibre.Authentication{}
		if err := json.Unmarshal(body, target); err != nil {
			t.Logf("Error unmarshalling Request Body: %s", err.Error())
			notAuthenticated(w)

			return
		}

		if target.Email != validEmail || target.Password != validPassword {
			t.Logf("Invalid email or password: E: %s, P: %s", target.Email, target.Password)
			notAuthenticated(w)

			return
		}

		validResponse, err := os.ReadFile("./test_files/client/response/login_successful.json")
		if err != nil {
			t.Logf("Error while serving valid response: %s", err.Error())
			notAuthenticated(w)

			return
		}

		w.Write(validResponse)
	})

	// Connections endpoint
	testServeMux.HandleFunc("/llu/connections", func(w http.ResponseWriter, r *http.Request) {
		// Validate that we have a valid JWT Token
		if r.Header.Get("Authorization") != validJWTToken {
			t.Logf("Authorization header was present, but JWT Token was invalid: %s", r.Header.Get("Authorization"))
			notAuthenticated(w)

			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Logf("Error reading Request Body: %s", err.Error())
			notAuthenticated(w)

			return
		}

		target := &golibre.Authentication{}
		if err := json.Unmarshal(body, target); err != nil {
			t.Logf("Error unmarshalling Request Body: %s", err.Error())
			notAuthenticated(w)

			return
		}

		if target.Email != validEmail || target.Password != validPassword {
			t.Logf("Invalid email or password: E: %s, P: %s", target.Email, target.Password)
			notAuthenticated(w)

			return
		}

		validResponse, err := os.ReadFile("./test_files/client/response/login_successful.json")
		if err != nil {
			t.Logf("Error while serving valid response: %s", err.Error())
			notAuthenticated(w)

			return
		}

		w.Write(validResponse)
	})

	// Add the handlers to an unstarted server
	srv := httptest.NewUnstartedServer(
		testServeMux,
	)

	// Set the TLS Config to skip the verify step, as this is a local test and outside the scope of the requirements
	srv.StartTLS()
	srv.TLS.InsecureSkipVerify = true

	// Return the server to the caller
	return srv, nil
}

func getTestServerAddress(testSrv *httptest.Server) string {
	url := testSrv.URL
	return strings.TrimPrefix(url, "https://")
}

func notAuthenticated(w http.ResponseWriter) {
	notAuthenticatedError := golibre.ErrorResponse{
		Status: 2,
		Error: golibre.ErrorMessage{
			Message: "notAuthenticated",
		},
	}

	notAuthenticatedErrorBytes, err := json.Marshal(notAuthenticatedError)
	if err != nil {
		panic(err)
	}

	// As of Go 1.17, Write() automatically sends http.StatusOK.
	w.Write(notAuthenticatedErrorBytes)
}

type RoundTripper struct {
	RoundTripFunc func(*http.Request) (*http.Response, error)
}

func (r RoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	return r.RoundTripFunc(request)
}

type RoundTripFunc func(t *testing.T, request *http.Request) (*http.Response, error)

func RoundTripperQueue(t *testing.T, queue []RoundTripFunc) http.RoundTripper {
	runNumber := 0

	return RoundTripper{
		RoundTripFunc: func(r *http.Request) (*http.Response, error) {
			defer func() {
				runNumber++
			}()

			if len(queue) <= runNumber {
				return nil, errors.New("empty queue")
			}

			return queue[runNumber](t, r)
		},
	}
}
