// Package bouncer TODO: docs
package bouncer

import (
	"context"
	"log"

	pb "github.com/prettyboiiii/bouncer/proto/bouncer"
	"google.golang.org/grpc"
)

// Instance TODO: docs
type Instance struct {
	pb.UnimplementedGreeterServer
	opt instanceOptions
}

// SayHello implements GreeterServer.
func (i *Instance) SayHello(_ context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", req.GetName())
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

// NewInstance TODO: docs
func NewInstance(server grpc.ServiceRegistrar, opts ...NewInstanceOption) *Instance {
	i := Instance{}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt.apply(&i.opt)
		}
	}

	pb.RegisterGreeterServer(server, &i)

	return &i
}
