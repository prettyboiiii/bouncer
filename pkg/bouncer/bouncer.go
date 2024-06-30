// Package bouncer implement essential types, methods for exporting
package bouncer

import (
	"google.golang.org/grpc"
)

type connection interface {
	grpc.ClientConnInterface
}

// Bouncer TODO: docs
type Bouncer struct {
	conn connection
}

// NewBouncer TODO: docs
func NewBouncer(conn connection) *Bouncer {
	return &Bouncer{
		conn: conn,
	}
}
