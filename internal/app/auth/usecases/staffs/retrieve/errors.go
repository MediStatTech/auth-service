package retrieve

import (
	errors "github.com/MediStatTech/error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToFindStaff = errors.NewGRPCError(codes.Internal, "failed to find staff")
)
