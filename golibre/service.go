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
	connectionService *ConnectionService
}

const defaultHTTPClientTimeout = time.Second * 15

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
				MinVersion: tls.TLS_AES_128_GCM_SHA256,
			},
		},
		timeout: defaultHTTPClientTimeout,
	}

	for _, opt := range opts {
		opt(config)
	}

	c := &client{
		httpClient: &http.Client{
			Transport: config.transport,
			Timeout:   config.timeout,
		},
		authentication: auth,
		apiURL:         apiURL,
		jwt: jwtAuth{
			mutex: &sync.Mutex{},
		},
	}

	return &Service{
		client: c,
		connectionService: &ConnectionService{
			client: c,
		},
	}
}

func (s *Service) GetMetadata(ctx context.Context) error {
	return nil
}

func (s *Service) Connection() *ConnectionService {
	return s.connectionService
}
