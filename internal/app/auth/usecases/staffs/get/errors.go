package get

import (
	errors "github.com/MediStatTech/error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToListStaffs = errors.NewGRPCError(codes.Internal, "failed to list staffs")
)
