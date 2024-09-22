package golibre_test

import (
	"context"
	"testing"

	"github.com/equalsgibson/golibre/golibre"
)

func TestConnection_GetData_200(t *testing.T) {
	t.Parallel()

	srv := newTestServer(t)
	defer srv.Close()

	testService := golibre.NewService(
		getTestServerAddress(srv),
		golibre.Authentication{
			Email:    validEmail,
			Password: validPassword,
		},
		golibre.WithTLSInsecureSkipVerify(),
	)

	ctx := context.Background()

	data, err := testService.Connection().GetConnectionData(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(data) != 1 {
		t.Fatal(data)
	}
}

func TestConnection_GetDataAndReuseJWTToken_200(t *testing.T) {
	t.Parallel()

	srv := newTestServer(t)
	defer srv.Close()

	testService := golibre.NewService(
		getTestServerAddress(srv),
		golibre.Authentication{
			Email:    validEmail,
			Password: validPassword,
		},
		golibre.WithTLSInsecureSkipVerify(),
	)

	ctx := context.Background()

	var err error

	_, err = testService.Connection().GetConnectionData(ctx)
	if err != nil {
		t.Fatal(err)
	}

	_, err = testService.Connection().GetConnectionData(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConnection_GetData_401(t *testing.T) {
	t.Parallel()

	srv := newTestServer(t)
	defer srv.Close()

	testService := golibre.NewService(
		getTestServerAddress(srv),
		golibre.Authentication{
			Email:    validEmail,
			Password: "invalidPassword",
		},
		golibre.WithTLSInsecureSkipVerify(),
	)

	ctx := context.Background()

	_, err := testService.Connection().GetConnectionData(ctx)
	if err == nil {
		t.Fatal("expected to get error due to bad password")
	}

	golibreErr, isCorrectErrorType := err.(*golibre.APIError)
	if !isCorrectErrorType {
		t.Fatal("did not get correct error type returned")
	}

	if golibreErr.Status != golibre.StatusUnauthenticated {
		t.Fatalf("expected to get golibre error with status code '%d', got status code '%d'", golibre.StatusUnauthenticated, golibreErr.Status)
	}

	if golibreErr.Detail.Message != "notAuthenticated" {
		t.Fatalf("expected to get different error message - expected: '%s', actual: '%s'", "notAuthenticated", golibreErr.Detail.Message)
	}
}
