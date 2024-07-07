// Package bouncer TODO: docs
package bouncer

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/prettyboiiii/bouncer/internal/syncx"
	pb "github.com/prettyboiiii/bouncer/proto/bouncer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var bouncersMap syncx.Map[string, *pb.BouncerData]

// Instance TODO: docs
type Instance struct {
	pb.UnimplementedBouncerServer
	opt instanceOptions
}

// GetAge implements bouncer.BouncerServer.
func (i *Instance) GetAge(ctx context.Context, req *pb.GetAgeRequest) (*durationpb.Duration, error) {
	if err := ctx.Err(); err != nil {
		return nil, status.FromContextError(err).Err()
	}
	b, found := bouncersMap.Load(req.BouncerId)
	if !found {
		return nil, BouncerNotFound
	}

	return durationpb.New(time.Since(b.CreatedAt.AsTime())), nil
}

// RegisterNewClient implements bouncer.BouncerServer.
func (i *Instance) RegisterNewClient(ctx context.Context, _ *pb.RegisterNewClientReq) (*pb.BouncerData, error) {
	if err := ctx.Err(); err != nil {
		return nil, status.FromContextError(err).Err()
	}
	b := pb.BouncerData{
		Id:        uuid.NewString(),
		CreatedAt: timestamppb.Now(),
	}
	bouncersMap.Store(b.Id, &b)

	i.opt.logger.Println("registering new client id:", b.Id)

	return &b, nil
}

// NewInstance TODO: docs
func NewInstance(server grpc.ServiceRegistrar, opts ...NewInstanceOption) *Instance {
	i := Instance{}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt.apply(&i.opt)
		}
	}

	applyDefaultOpt(&i.opt)

	pb.RegisterBouncerServer(server, &i)

	return &i
}

func applyDefaultOpt(io *instanceOptions) {
	defaultOpt := getDefaultOpt()
	if io.name == "" {
		io.name = defaultOpt.name
	}
	if io.logger == nil {
		io.logger = defaultOpt.logger
	}
}
