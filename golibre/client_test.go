package golibre_test

import (
	"context"
	"testing"

	"github.com/equalsgibson/golibre/golibre"
)

func TestClientLogin_200(t *testing.T) {
	srv, err := newTestServer(t)
	if err != nil {
		t.Fatal(err)
	}
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

	t.Fatal(data)
}
