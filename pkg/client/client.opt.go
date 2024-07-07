package client

import (
	"google.golang.org/grpc"
)

type clientOptions struct {
	gRPCOpts []grpc.DialOption
}

// Option TODO: docs
type Option interface {
	apply(*clientOptions)
}

type funcDialOption struct {
	f func(*clientOptions)
}

func (f *funcDialOption) apply(io *clientOptions) {
	f.f(io)
}

// GRPCWrapper TODO: docs
func GRPCWrapper(opts ...grpc.DialOption) Option {
	return newFuncClientOption(func(co *clientOptions) {
		co.gRPCOpts = append(co.gRPCOpts, opts...)
	})
}

func newFuncClientOption(f func(*clientOptions)) *funcDialOption {
	return &funcDialOption{
		f: f,
	}
}
