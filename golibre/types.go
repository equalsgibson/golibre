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

type ErrorResponse struct {
	Status uint         `json:"status"`
	Error  ErrorMessage `json:"error"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type ConnectionResponse struct {
	Status uint             `json:"status"`
	Data   []ConnectionData `json:"data"`
	Ticket AuthTicket       `json:"ticket"`
}

type ConnectionData struct {
	ID                 UserID             `json:"id"`
	PatientID          PatientID          `json:"patientId"`
	Country            string             `json:"country"`
	Status             uint               `json:"status"`
	FirstName          string             `json:"firstName"`
	LastName           string             `json:"lastName"`
	TargetLow          uint               `json:"targetLow"`
	TargetHigh         uint               `json:"targetHigh"`
	UnitOfMeasurement  UnitOfMeasurement  `json:"uom"`
	Sensor             Sensor             `json:"sensor"`
	AlarmRules         AlarmRules         `json:"alarmRules"`
	GlucoseMeasurement GlucoseMeasurement `json:"glucoseMeasurement"`
	GlucoseItem        GlucoseMeasurement `json:"glucoseItem"`
	GlucoseAlarm       any                `json:"glucoseAlarm"`
	PatientDevice      PatientDevice      `json:"patientDevice"`
	Created            uint               `json:"created"`
}

type PatientDevice struct {
	DID                 string              `json:"did"`
	DTID                uint                `json:"dtid"`
	V                   string              `json:"v"`
	LL                  uint                `json:"ll"`
	HL                  uint                `json:"hl"`
	U                   uint                `json:"u"`
	FixedLowAlarmValues FixedLowAlarmValues `json:"fixedLowAlarmValues"`
	Alarms              bool                `json:"alarms"`
	FixedLowThreshold   uint                `json:"fixedLowThreshold"`
}

type FixedLowAlarmValues struct {
	MGDL  uint    `json:"mgdl"`
	MMOLL float32 `json:"mmoll"`
}

type GlucoseMeasurement struct {
	FactoryTimestamp Timestamp `json:"FactoryTimestamp"` //nolint:tagliatelle
	Timestamp        Timestamp `json:"Timestamp"`        //nolint:tagliatelle
	Type             uint      `json:"type"`
	ValueInMgPerDl   uint      `json:"ValueInMgPerDl"`   //nolint:tagliatelle
	TrendArrow       uint      `json:"TrendArrow"`       //nolint:tagliatelle
	MeasurementColor uint      `json:"MeasurementColor"` //nolint:tagliatelle
	GlucoseUnits     uint      `json:"GlucoseUnits"`     //nolint:tagliatelle
	Value            float32   `json:"Value"`            //nolint:tagliatelle
	IsHigh           bool      `json:"isHigh"`
	IsLow            bool      `json:"isLow"`
}

type Sensor struct {
	DeviceID     string `json:"deviceId"`
	SerialNumber string `json:"sn"`
	Activated    uint   `json:"a"`
	W            uint   `json:"w"`
	PT           uint   `json:"pt"`
	S            bool   `json:"s"`
	LJ           bool   `json:"lj"`
}

type AlarmRules struct {
	C   bool        `json:"c"`
	H   AlarmRuleH  `json:"h"`
	F   AlarmRule   `json:"f"`
	L   AlarmRule   `json:"l"`
	ND  AlarmRuleND `json:"nd"`
	P   uint        `json:"p"`
	R   uint        `json:"r"`
	STD any         `json:"std"`
}

type AlarmRuleH struct {
	TH   uint    `json:"th"`
	THMM float32 `json:"thmm"`
	D    uint    `json:"d"`
	F    float32 `json:"f"`
}

type AlarmRule struct {
	TH   uint    `json:"th"`
	THMM float32 `json:"thmm"`
	D    uint    `json:"d"`
	TL   uint    `json:"tl"`
	TLMM uint    `json:"tlmm"`
}

type AlarmRuleND struct {
	I uint `json:"i"`
	R uint `json:"r"`
	L uint `json:"l"`
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
