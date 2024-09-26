package golibre_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/equalsgibson/golibre/golibre"
)

const (
	validPassword string = "VALID_PASSWORD"
	validEmail    string = "EMAIL"
	validJWTToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4LTEyMzQtMTIzNC0xMjM0LTQ2MDRhNjdiN2FiNCIsImZpcnN0TmFtZSI6IlNvbWUiLCJsYXN0TmFtZSI6Ik9uZSIsImNvdW50cnkiOiJHQiIsInJlZ2lvbiI6ImV1MiIsInJvbGUiOiJwYXRpZW50IiwidW5pdHMiOjAsInByYWN0aWNlcyI6W10sImMiOjEsInMiOiJsbHUuYW5kcm9pZCIsImV4cCI6MTc0MjI5NDIwN30.ilRwCINRf6nQViQ9c0BLZD9x21qsiBx43EzMk1POTuk" // #nosec G101 Fake Token For Test

	validPatientID golibre.PatientID = "87654321-4321-4321-4321-0242ac110002"
)

func authenticatedRequestMiddleware(t *testing.T, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer "+validJWTToken {
			t.Logf("Authorization header was present, but JWT Token was invalid: %s", r.Header.Get("Authorization"))
			notAuthenticated(w)

			return
		}

		next.ServeHTTP(w, r)
	})
}

func newTestServer(t *testing.T) *httptest.Server {
	t.Helper()

	testServeMux := http.NewServeMux()

	// Private Endpoints
	// -> User endpoint
	testServeMux.Handle("/user", authenticatedRequestMiddleware(t, userHandler(t)))
	// -> Account endpoint
	testServeMux.Handle("/account", authenticatedRequestMiddleware(t, accountHandler(t)))
	// -> Connections endpoint
	testServeMux.Handle("/llu/connections", authenticatedRequestMiddleware(t, connectionHandler(t)))
	// -> ConnectionGraph endpoint
	testServeMux.Handle("/llu/connections/{patientID}/graph", authenticatedRequestMiddleware(t, connectionGraphHandler(t)))

	// Public Endpoints
	// -> Authentication endpoint
	testServeMux.HandleFunc("/llu/auth/login", loginHandler(t))
	testServeMux.HandleFunc("/", http.NotFound)

	// Add the handlers to an unstarted server
	srv := httptest.NewUnstartedServer(
		testServeMux,
	)

	// Set the TLS Config to skip the verify step, as this is a local test and outside the scope of the requirements
	srv.StartTLS()
	srv.TLS.InsecureSkipVerify = true

	// Return the server to the caller
	return srv
}

func getTestServerAddress(testSrv *httptest.Server) string {
	url := testSrv.URL

	return strings.TrimPrefix(url, "https://")
}

func notAuthenticated(w http.ResponseWriter) {
	response, err := os.ReadFile("./test_files/error/response/error_unauthenticated.json")
	if err != nil {
		panic(err)
	}

	// As of Go 1.17, Write() automatically sends http.StatusOK.
	_, err = w.Write(response)
	if err != nil {
		panic(err)
	}
}

func userHandler(t *testing.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		validResponse, err := os.ReadFile("./test_files/user/response/getUserData_200.json")
		if err != nil {
			t.Logf("Error while reading response file: %s", err.Error())

			panic(err)
		}

		_, err = w.Write(validResponse)
		if err != nil {
			t.Logf("Error while writing response bytes: %s", err.Error())

			panic(err)
		}
	})
}

func accountHandler(t *testing.T) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		validResponse, err := os.ReadFile("./test_files/account/response/getAccountDetails_200.json")
		if err != nil {
			t.Logf("Error while serving valid response: %s", err.Error())

			panic(err)
		}

		_, err = w.Write(validResponse)
		if err != nil {
			panic(err)
		}
	})
}

func loginHandler(t *testing.T) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		_, err = w.Write(validResponse)
		if err != nil {
			panic(err)
		}
	})
}

func connectionHandler(t *testing.T) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		validResponse, err := os.ReadFile("./test_files/connection/response/getAllConnectionData_200.json")
		if err != nil {
			t.Logf("Error while serving valid response: %s", err.Error())
			notAuthenticated(w)

			return
		}

		_, err = w.Write(validResponse)
		if err != nil {
			panic(err)
		}
	})
}

func connectionGraphHandler(t *testing.T) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.PathValue("patientID") != string(validPatientID) {
			w.WriteHeader(http.StatusMethodNotAllowed)

			return
		}

		validResponse, err := os.ReadFile("./test_files/connection/response/getConnectionGraph_200.json")
		if err != nil {
			t.Logf("Error while serving valid response: %s", err.Error())
			notAuthenticated(w)

			return
		}

		_, err = w.Write(validResponse)
		if err != nil {
			panic(err)
		}
	})
}
