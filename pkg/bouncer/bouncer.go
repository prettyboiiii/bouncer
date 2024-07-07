// Package bouncer implement essential types, methods for exporting
package bouncer

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/prettyboiiii/bouncer/proto/bouncer"
	"google.golang.org/grpc"
)

type connection interface {
	grpc.ClientConnInterface
}

// Bouncer TODO: docs
type Bouncer struct {
	*pb.BouncerData
	conn connection
}

// Age TODO: docs
func (b *Bouncer) Age() time.Duration {
	return time.Since(time.Now())
}

// NewBouncer TODO: docs
func NewBouncer(ctx context.Context, conn connection) *Bouncer {
	client := pb.NewBouncerClient(conn)
	res, err := client.RegisterNewClient(ctx, &pb.RegisterNewClientReq{})
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(res)

	b := Bouncer{
		conn:        conn,
		BouncerData: res,
	}

	for {
		time.Sleep(1 * time.Second)
		log.Println(client.GetAge(context.TODO(), &pb.GetAgeRequest{
			BouncerId: b.Id,
		}))
	}

	// return &b
}
