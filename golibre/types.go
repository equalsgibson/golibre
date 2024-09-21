package golibre

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
	ID                 UserID            `json:"id"`
	PatientID          PatientID         `json:"patientId"`
	Country            string            `json:"country"`
	Status             uint              `json:"status"`
	FirstName          string            `json:"firstName"`
	LastName           string            `json:"lastName"`
	TargetLow          uint              `json:"targetLow"`
	TargetHigh         uint              `json:"targetHigh"`
	UnitOfMeasurement  UnitOfMeasurement `json:"uom"`
	Sensor             Sensor            `json:"sensor"`
	AlarmRules         AlarmRules        `json:"alarmRules"`
	GlucoseMeasurement any
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

type PatientID string
type UserID string
type UnitOfMeasurement uint

const (
	MMOL UnitOfMeasurement = iota
	DL
)
