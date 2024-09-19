package golibre

import "context"

type ConnectionService struct {
	client *client
}

func (c *ConnectionService) GetConnectionData(ctx context.Context) error {
	return nil
}
