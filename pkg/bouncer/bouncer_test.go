// nolint:unused,revive
package bouncer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

type mockConn struct{}

// Invoke implements connection.
func (m *mockConn) Invoke(_ context.Context, method string, args any, reply any, opts ...grpc.CallOption) error {
	panic("unimplemented")
}

// NewStream implements connection.
func (m *mockConn) NewStream(
	_ context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	panic("unimplemented")
}

type newBouncerTest struct {
	name   string
	expect *Bouncer
}

func TestNewBouncer(t *testing.T) {
	tests := []newBouncerTest{
		{
			name:   "normal",
			expect: NewBouncer(&mockConn{}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, NewBouncer(&mockConn{}))
		})
	}
}
