package golibre

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Service struct {
	client            *client
	connectionService *ConnectionService
}

const DefaultHTTPClientTimeout = time.Second * 15

func NewService(
	subDomain string,
	auth Authentication,
	opts ...configOption,
) *Service {
	config := &config{
		userAgent: "equalsgibson/golibre",
		timeout:   DefaultHTTPClientTimeout,
	}

	for _, opt := range opts {
		opt(config)
	}

	c := &client{
		httpClient: &http.Client{
			Transport: config.roundTripper,
		},
		userAgent: config.userAgent,
		subDomain: subDomain,
		jwt: jwtAuth{
			mutex: &sync.Mutex{},
		},
	}

	return &Service{
		client: c,
	}
}

func (s *Service) GetMetadata(ctx context.Context) error {
	return nil
}

func (s *Service) Connection() *ConnectionService {
	return s.connectionService
}
