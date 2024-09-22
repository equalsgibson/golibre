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

	t.Fatal("got to end   ", data)
}
