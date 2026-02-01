package jwt

import (
	errors "github.com/MediStatTech/error"
	"google.golang.org/grpc/codes"
)

var (
	errRequestNil = errors.NewGRPCError(codes.InvalidArgument, "request is nil")
	errTokenRequired = errors.NewGRPCError(codes.InvalidArgument, "token is required")
)