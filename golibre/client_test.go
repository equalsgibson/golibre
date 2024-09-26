package golibre_test

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"testing"

	"github.com/equalsgibson/golibre/golibre"
)

func TestClientLogin_200(t *testing.T) {
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

	_, err := testService.Connection().GetAllConnectionData(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRequestPreProcessor(t *testing.T) {
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
		golibre.WithRequestPreProcessor(golibre.RequestPreProcessorFunc(
			func(r *http.Request) error {
				return nil
			},
		)),
		golibre.WithSlogger(slog.New(slog.NewJSONHandler(io.Discard, nil))),
	)

	ctx := context.Background()

	_, err := testService.Connection().GetAllConnectionData(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
