package golibre

import (
	"encoding/json"
	"time"
)

type LoginResponse struct {
	Status uint      `json:"status"`
	Data   LoginData `json:"data"`
}

type LoginData struct {
	User               map[string]any `json:"user"`
	Messages           map[string]any `json:"messages"`
	Notifications      map[string]any `json:"notifications"`
	AuthTicket         AuthTicket     `json:"authTicket"`
	Invitations        []any          `json:"invitations"`
	TrustedDeviceToken string         `json:"trustedDeviceToken"`
}

type AuthTicket struct {
	Token    string `json:"token"`
	Expires  uint64 `json:"expires"`
	Duration uint64 `json:"duration"`
}

type (
	PatientID         string
	UserID            string
	UnitOfMeasurement uint
)

const (
	MMOL UnitOfMeasurement = iota
	DL
)

/*
Timestamp requires custom unmarshalling due to being returned
by the Libreview API in the following format: 'M/D/YYYY H:M:S PM'

Example data returned: "9/20/2024 3:42:05 PM".
*/
type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	time, err := time.Parse("1/2/2006 3:4:5 PM", s)
	if err != nil {
		return err
	}

	*t = Timestamp(time)

	return nil
}

type StatusCode uint

const (
	StatusOK StatusCode = iota
	_
	StatusUnauthenticated
)
