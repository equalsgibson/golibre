package golibre_test

import (
	"context"
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
