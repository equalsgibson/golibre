package golibre_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/equalsgibson/golibre"
)

func TestConnection_GetAllConnections_200(t *testing.T) {
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

	actual, err := testService.Connection().GetAllConnectionData(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(actual) != 1 {
		t.Fatal(actual)
	}
}

func TestConnection_GetAllConnectionsAndReuseJWTToken_200(t *testing.T) {
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

	_, err = testService.Connection().GetAllConnectionData(ctx)
	if err != nil {
		t.Fatal(err)
	}

	_, err = testService.Connection().GetAllConnectionData(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConnection_GetAllConnections_Unauthenticated(t *testing.T) {
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

	_, err := testService.Connection().GetAllConnectionData(ctx)
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

func TestConnection_GetConnectionGraph_200(t *testing.T) {
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

	factoryTimestamp, err := time.Parse("1/2/2006 3:4:5 PM", "9/22/2024 8:51:13 AM")
	if err != nil {
		t.Fatal(err)
	}

	expected := []golibre.GraphGlucoseMeasurement{
		{
			FactoryTimestamp: golibre.Timestamp(factoryTimestamp),
			Timestamp:        golibre.Timestamp(factoryTimestamp.Add(time.Hour)),
			Type:             0,
			ValueInMgPerDl:   193,
			MeasurementColor: 2,
			GlucoseUnits:     0,
			Value:            10.7,
			IsHigh:           false,
			IsLow:            false,
		},
		{
			FactoryTimestamp: golibre.Timestamp(factoryTimestamp.Add(time.Minute * 5)),
			Timestamp:        golibre.Timestamp(factoryTimestamp.Add(time.Hour).Add(time.Minute * 5)),
			Type:             0,
			ValueInMgPerDl:   195,
			MeasurementColor: 2,
			GlucoseUnits:     0,
			Value:            10.8,
			IsHigh:           false,
			IsLow:            false,
		},
		{
			FactoryTimestamp: golibre.Timestamp(factoryTimestamp.Add(time.Minute * 10)),
			Timestamp:        golibre.Timestamp(factoryTimestamp.Add(time.Hour).Add(time.Minute * 10)),
			Type:             0,
			ValueInMgPerDl:   202,
			MeasurementColor: 2,
			GlucoseUnits:     0,
			Value:            11.2,
			IsHigh:           false,
			IsLow:            false,
		},
		{
			FactoryTimestamp: golibre.Timestamp(factoryTimestamp.Add(time.Minute * 15)),
			Timestamp:        golibre.Timestamp(factoryTimestamp.Add(time.Hour).Add(time.Minute * 15)),
			Type:             0,
			ValueInMgPerDl:   205,
			MeasurementColor: 2,
			GlucoseUnits:     0,
			Value:            11.4,
			IsHigh:           false,
			IsLow:            false,
		},
	}

	actual, err := testService.Connection().GetConnectionGraph(ctx, validPatientID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual.GraphData, expected) {
		t.Log("did not receive the expected values")
		t.Fatalf("\n\nEXPECTED: %+v\n\nACTUAL: %+v\n\n", expected, actual.GraphData)
	}
}
