package conn

import "google.golang.org/grpc"

type connOption interface {
	grpc.DialOption
}
