package sign_in

import (
	errors "github.com/MediStatTech/error"
	"google.golang.org/grpc/codes"
)

var (
	errInvalidCredentials     = errors.NewGRPCError(codes.Unauthenticated, "invalid credentials")
	errStaffInactive          = errors.NewGRPCError(codes.Unauthenticated, "staff is inactive")
	errFailedToGenerateJWT     = errors.NewGRPCError(codes.Internal, "failed to generate JWT")
)
