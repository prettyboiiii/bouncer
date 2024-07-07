package bouncer

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// BouncerNotFound TODO: docs
var BouncerNotFound = status.Error(codes.NotFound, "bouncer not found")
