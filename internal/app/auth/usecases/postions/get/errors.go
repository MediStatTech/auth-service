package get

import (
	errors "github.com/MediStatTech/error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToListPositions = errors.NewGRPCError(codes.Internal, "failed to list positions")
)
