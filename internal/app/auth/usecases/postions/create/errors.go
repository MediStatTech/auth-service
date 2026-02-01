package create

import (
	errors "github.com/MediStatTech/error"
	"google.golang.org/grpc/codes"
)

var (
	errPositionAlreadyExists   = errors.NewGRPCError(codes.AlreadyExists, "position already exists")
	errFailedToCreatePosition  = errors.NewGRPCError(codes.Internal, "failed to create position")
)
