package golibre

import (
	"context"
	"net/http"
)

type AccountService struct {
	client *client
}

func (a *AccountService) GetAccountDetails(ctx context.Context) (AccountDetailData, error) {
	endpoint := "/account"

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		endpoint,
		http.NoBody,
	)
	if err != nil {
		return AccountDetailData{}, err
	}

	target := AccountDetailsResponse{}
	if err := a.client.Do(req, &target); err != nil {
		return AccountDetailData{}, err
	}

	return target.Data, nil
}

type AccountDetailsResponse BaseResponse[AccountDetailData]

type AccountDetailData struct {
	User UserAccountData `json:"user"`
}

type UserAccountData struct {
	ID                    UserID         `json:"id"`
	FirstName             string         `json:"firstName"`
	LastName              string         `json:"lastName"`
	DateOfBirth           uint           `json:"dateOfBirth"`
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
}
