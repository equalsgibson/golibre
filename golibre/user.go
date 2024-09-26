package golibre

import (
	"context"
	"net/http"
)

type UserService struct {
	client *client
}

func (u *UserService) GetLoggedInUser(ctx context.Context) (UserData, error) {
	endpoint := "/user"

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		endpoint,
		http.NoBody,
	)
	if err != nil {
		return UserData{}, err
	}

	target := GetLoggedInUserResponse{}
	if err := u.client.Do(req, &target); err != nil {
		return UserData{}, err
	}

	return target.Data, nil
}

type GetLoggedInUserResponse BaseResponse[UserData]

type UserData struct {
	User               User          `json:"user"`
	Messages           Messages      `json:"messages"`
	Notifications      Notifications `json:"notifications"`
	AuthTicket         AuthTicket    `json:"authTicket"`
	Invitations        []string      `json:"invitations"`
	TrustedDeviceToken string        `json:"trustedDeviceToken"`
}

type User struct {
	ID                    UserID         `json:"id"`
	FirstName             string         `json:"firstName"`
	LastName              string         `json:"lastName"`
	Email                 string         `json:"email"`
	Country               string         `json:"country"`
	UILanguage            string         `json:"uiLanguage"`
	CommunicationLanguage string         `json:"communicationLanguage"`
	AccountType           string         `json:"accountType"`
	UOM                   string         `json:"uom"`
	DateFormat            string         `json:"dateFormat"`
	TimeFormat            string         `json:"timeFormat"`
	EmailDay              []uint         `json:"emailDay"`
	System                System         `json:"system"`
	Details               map[string]any `json:"details"`
	TwoFactor             TwoFactor      `json:"twoFactor"`
	Created               uint           `json:"created"`
	LastLogin             uint           `json:"lastLogin"`
	Programs              map[string]any `json:"programs"`
	DateOfBirth           uint           `json:"dateOfBirth"`
	Devices               map[string]any `json:"devices"`
	Consents              Consents       `json:"consents"`
}

type Messages struct {
	Unread uint `json:"unread"`
}

type Notifications struct {
	Unresolved uint `json:"unresolved"`
}

type System struct {
	Messages SystemMessages `json:"messages"`
}

type SystemMessages struct {
	AppReviewBanner                  uint   `json:"appReviewBanner"`
	FirstUsePhoenix                  uint   `json:"firstUsePhoenix"`
	FirstUsePhoenixReportsDataMerged uint   `json:"firstUsePhoenixReportsDataMerged"`
	LLUGettingStartedBanner          uint   `json:"lluGettingStartedBanner"`
	LLUNewFeatureModal               uint   `json:"lluNewFeatureModal"`
	LLUOnboarding                    uint   `json:"lluOnboarding"`
	LVWebPostRelease                 string `json:"lvWebPostRelease"`
}

type TwoFactor struct {
	PrimaryMethod   string `json:"primaryMethod"`
	PrimaryValue    string `json:"primaryValue"`
	SecondaryMethod string `json:"secondaryMethod"`
	SecondaryValue  string `json:"secondaryValue"`
}

type Consents struct {
	LLU               LLUConsent               `json:"llu"`
	RealWorldEvidence RealWorldEvidenceConsent `json:"realWorldEvidence"`
}

type LLUConsent struct {
	PolicyAccept uint `json:"policyAccept"`
	TOUAccept    uint `json:"touAccept"`
}

type RealWorldEvidenceConsent struct {
	PolicyAccept uint           `json:"policyAccept"`
	TOUAccept    uint           `json:"touAccept"`
	History      []PolicyAccept `json:"history"`
}

type PolicyAccept struct {
	PolicyAccept uint `json:"policyAccept"`
}
