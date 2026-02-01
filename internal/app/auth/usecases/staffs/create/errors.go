package create

import (
	errors "github.com/MediStatTech/error"
	"google.golang.org/grpc/codes"
)

var (
	errEmailAlreadyExists   = errors.NewGRPCError(codes.AlreadyExists, "email already exists")
	errFailedToHashPassword = errors.NewGRPCError(codes.Internal, "failed to hash password")
	errFailedToCreateStaff  = errors.NewGRPCError(codes.Internal, "failed to create staff")
)
