package staffs

import (
	errors "github.com/MediStatTech/error"
	"google.golang.org/grpc/codes"
)

var (
	errRequestNil         = errors.NewGRPCError(codes.InvalidArgument, "request is nil")
	errEmailRequired      = errors.NewGRPCError(codes.InvalidArgument, "email is required")
	errPasswordRequired   = errors.NewGRPCError(codes.InvalidArgument, "password is required")
	errPositionIDRequired = errors.NewGRPCError(codes.InvalidArgument, "position ID is required")
	errFirstNameRequired  = errors.NewGRPCError(codes.InvalidArgument, "first name is required")
	errLastNameRequired   = errors.NewGRPCError(codes.InvalidArgument, "last name is required")
	errStaffIDRequired    = errors.NewGRPCError(codes.InvalidArgument, "staff ID is required")
)
