// Package conn TODO: docs
package conn

import (
	"fmt"

	"google.golang.org/grpc"
)

// Conn manages connections to the bouncer server.
type Conn struct {
	connection *grpc.ClientConn
}

// Close closes the connection to the bouncer server.
func (c *Conn) Close() error {
	if c.connection != nil {
		err := c.connection.Close()
		if err != nil {
			return fmt.Errorf("failed to close connection: %v", err)
		}
		c.connection = nil
	}
	return nil
}

// NewConn creates a new ConnManager with the given server address and timeout.
func NewConn(address string, opts ...connOption) (*Conn, error) {
	grpcOpts := make([]grpc.DialOption, len(opts))
	for i, o := range opts {
		grpcOpts[i] = grpc.DialOption(o)
	}
	c, err := grpc.NewClient(address, grpcOpts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection: %v", err)
	}

	return &Conn{
		connection: c,
	}, nil
}
