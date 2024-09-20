package golibre

import (
	"context"
	"net/http"
)

type ConnectionService struct {
	client *client
}

func (c *ConnectionService) GetConnectionData(ctx context.Context) (LoginResponse, error) {
	endpoint := "/llu/connections"

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		endpoint,
		http.NoBody,
	)
	if err != nil {
		return LoginResponse{}, err
	}

	target := LoginResponse{}
	if err := c.client.Do(req, &target); err != nil {
		return LoginResponse{}, err
	}

	return target, nil
}
