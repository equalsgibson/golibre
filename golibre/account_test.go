package golibre_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/equalsgibson/golibre/golibre"
)

func TestConnection_GetAccountDetails_200(t *testing.T) {
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

	expected := golibre.AccountDetailData{
		User: golibre.UserAccountData{
			ID:                    "12345678-1234-1234-1234-4604a67b7ab4",
			FirstName:             "Some",
			LastName:              "One",
			DateOfBirth:           123456,
			Email:                 "someone+libre@gmail.com",
			Country:               "GB",
			UILanguage:            "en-US",
			CommunicationLanguage: "en-US",
			AccountType:           "pat",
			UOM:                   "0",
			DateFormat:            "2",
			TimeFormat:            "2",
			EmailDay:              []uint{1},
			System: golibre.System{
				Messages: golibre.SystemMessages{
					AppReviewBanner:                  1720130314,
					FirstUsePhoenix:                  1720130262,
					FirstUsePhoenixReportsDataMerged: 1720130262,
					LLUGettingStartedBanner:          1720130318,
					LLUNewFeatureModal:               1720130302,
					LLUOnboarding:                    1720130310,
					LVWebPostRelease:                 "3.18.15",
				},
			},
			Details: map[string]any{},
			TwoFactor: golibre.TwoFactor{
				PrimaryMethod:   "phone",
				PrimaryValue:    "+44 1234567891",
				SecondaryMethod: "email",
				SecondaryValue:  "someone+libre@gmail.com",
			},
			Created:   1720130262,
			LastLogin: 1727345530,
		},
	}

	actual, err := testService.Account().GetAccountDetails(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Log("did not receive the expected values")
		t.Fatalf("\n\nEXPECTED: %+v\n\nACTUAL: %+v\n\n", expected, actual)
	}
}
