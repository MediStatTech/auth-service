package positions

import (
	errors "github.com/MediStatTech/error"
	"google.golang.org/grpc/codes"
)

var (
	errRequestNil    = errors.NewGRPCError(codes.InvalidArgument, "request is nil")
	errNameRequired = errors.NewGRPCError(codes.InvalidArgument, "name is required")
)