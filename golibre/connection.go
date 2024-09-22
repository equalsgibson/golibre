package golibre

import (
	"context"
	"net/http"
)

type ConnectionService struct {
	client *client
}

func (c *ConnectionService) GetConnectionData(ctx context.Context) ([]ConnectionData, error) {
	endpoint := "/llu/connections"

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		endpoint,
		http.NoBody,
	)
	if err != nil {
		return nil, err
	}

	target := ConnectionResponse{}
	if err := c.client.Do(req, &target); err != nil {
		return nil, err
	}

	return target.Data, nil
}
