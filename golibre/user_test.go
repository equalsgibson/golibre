package golibre_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/equalsgibson/golibre/golibre"
)

func TestConnection_GetLoggedInUser_200(t *testing.T) {
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

	expected := golibre.UserData{
		User: golibre.User{
			ID:                    "12345678-1234-1234-1234-4604a67b7ab4",
			FirstName:             "Some",
			LastName:              "One",
			Email:                 "someone+libre@gmail.com",
			DateOfBirth:           123456,
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
			Programs:  map[string]any{},
			Devices:   map[string]any{},
			Consents: golibre.Consents{
				LLU: golibre.LLUConsent{
					PolicyAccept: 1720130262,
					TOUAccept:    1727258891,
				},
				RealWorldEvidence: golibre.RealWorldEvidenceConsent{
					PolicyAccept: 1727044734,
					TOUAccept:    0,
					History: []golibre.PolicyAccept{
						{
							PolicyAccept: 1727044734,
						},
					},
				},
			},
		},
		Messages: golibre.Messages{
			Unread: 0,
		},
		Notifications: golibre.Notifications{
			Unresolved: 0,
		},
		AuthTicket: golibre.AuthTicket{
			Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEyMzQ1Njc4LTEyMzQtMTIzNC0xMjM0LTQ2MDRhNjdiN2FiNCIsImZpcnN0TmFtZSI6IlNvbWUiLCJsYXN0TmFtZSI6Ik9uZSIsImNvdW50cnkiOiJHQiIsInJlZ2lvbiI6ImV1MiIsInJvbGUiOiJwYXRpZW50IiwidW5pdHMiOjAsInByYWN0aWNlcyI6W10sImMiOjEsInMiOiJsbHUuYW5kcm9pZCIsImV4cCI6MTc0MjI5NDIwN30.ilRwCINRf6nQViQ9c0BLZD9x21qsiBx43EzMk1POTuk",
			Expires:  1742294207,
			Duration: 15552000000,
		},
		Invitations: []string{
			"abcdefg-abcd-abcd-abcd-a20cf1532aef",
		},
		TrustedDeviceToken: "",
	}

	actual, err := testService.User().GetLoggedInUser(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Log("did not receive the expected values")
		t.Fatalf("\n\nEXPECTED: %+v\n\nACTUAL: %+v\n\n", expected, actual)
	}
}
