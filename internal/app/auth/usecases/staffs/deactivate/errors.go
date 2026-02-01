package deactivate

import (
	errors "github.com/MediStatTech/error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToFindStaff   = errors.NewGRPCError(codes.NotFound, "failed to find staff")
	errFailedToUpdateStaff = errors.NewGRPCError(codes.Internal, "failed to update staff")
)
