// Package client TODO: docs
package client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

// BouncerClient manages connections to the bouncer server.
type BouncerClient struct {
	connection *grpc.ClientConn
	opts       clientOptions
}

// Invoke sends the RPC request on the wire and returns after response is
// received.  This is typically called by generated code.
//
// All errors returned by Invoke are compatible with the status package.
func (c *BouncerClient) Invoke(ctx context.Context, method string, args any, reply any, opts ...grpc.CallOption) error {
	return c.connection.Invoke(ctx, method, args, reply, opts...)
}

// NewStream creates a new Stream for the client side. This is typically
// called by generated code. ctx is used for the lifetime of the stream.
//
// To ensure resources are not leaked due to the stream returned, one of the following
// actions must be performed:
//
//  1. Call Close on the ClientConn.
//  2. Cancel the context provided.
//  3. Call RecvMsg until a non-nil error is returned. A protobuf-generated
//     client-streaming RPC, for instance, might use the helper function
//     CloseAndRecv (note that CloseSend does not Recv, therefore is not
//     guaranteed to release all resources).
//  4. Receive a non-nil, non-io.EOF error from Header or SendMsg.
//
// If none of the above happen, a goroutine and a context will be leaked, and grpc
// will not call the optionally-configured stats handler with a stats.End message.
func (c *BouncerClient) NewStream(
	ctx context.Context,
	desc *grpc.StreamDesc,
	method string,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	return c.connection.NewStream(ctx, desc, method, opts...)
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
