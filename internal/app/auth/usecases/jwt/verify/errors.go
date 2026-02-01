package verify

import (
	errors "github.com/MediStatTech/error"
	"google.golang.org/grpc/codes"
)

var (
	errFailedToVerifyJWT = errors.NewGRPCError(codes.Internal, "failed to verify JWT")
)