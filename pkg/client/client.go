// Package client TODO: docs
package client

import (
	"fmt"

	"google.golang.org/grpc"
)

// BouncerClient manages connections to the bouncer server.
type BouncerClient struct {
	connection *grpc.ClientConn
	opts       clientOptions
}

// Close closes the connection to the bouncer server.
func (c *BouncerClient) Close() error {
	if c.connection != nil {
		err := c.connection.Close()
		if err != nil {
			return fmt.Errorf("failed to close connection: %v", err)
		}
		c.connection = nil
	}
	return nil
}

// Dial creates a new BouncerClient with the given server address.
func Dial(address string, opts ...Option) (*BouncerClient, error) {
	bc := BouncerClient{}
	for _, o := range opts {
		o.apply(&bc.opts)
	}

	c, err := grpc.NewClient(address, bc.opts.gRPCOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection: %v", err)
	}
	bc.connection = c

	return &bc, nil
}
