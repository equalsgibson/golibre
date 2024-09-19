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
	Invitations        map[string]any `json:"invitations"`
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
