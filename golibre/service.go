package golibre

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"sync"
	"time"
)

type Service struct {
	client            *client
	accountService    *AccountService
	connectionService *ConnectionService
	userService       *UserService
}

func NewService(
	apiURL string,
	auth Authentication,
	opts ...configOption,
) *Service {
	config := &config{
		transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext(ctx, network, addr)
			},
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS13,
			},
		},
	}

	for _, opt := range opts {
		opt(config)
	}

	c := &client{
		httpClient: &http.Client{
			Transport: config.transport,
			Timeout:   15 * time.Second,
		},
		authentication: auth,
		apiURL:         apiURL,
		jwt: jwtAuth{
			mutex: &sync.RWMutex{},
		},
	}

	return &Service{
		client: c,
		accountService: &AccountService{
			client: c,
		},
		connectionService: &ConnectionService{
			client: c,
		},
		userService: &UserService{
			client: c,
		},
	}
}

func (s *Service) Connection() *ConnectionService {
	return s.connectionService
}

func (s *Service) User() *UserService {
	return s.userService
}

func (s *Service) Account() *AccountService {
	return s.accountService
}
