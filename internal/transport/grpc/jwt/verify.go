package jwt

import (
	"context"

	auth_pb "github.com/MediStatTech/auth-client/pb/go/services/v1"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/jwt/verify"
)

func (h *Handler) VerifyToken(
	ctx context.Context,
	req *auth_pb.VerifyTokenRequest,
) (*auth_pb.VerifyTokenReply, error) {
	if err := validateVerifyJWTRequest(req); err != nil {
		return nil, err
	}

	reply, err := h.commands.Verify.Execute(ctx, verify.Request{
		Token: req.Token,
	})
	if err != nil {
		return nil, err
	}

	return &auth_pb.VerifyTokenReply{
		StaffId:    reply.StaffID,
		PositionId: reply.PositionID,
	}, nil
}

func validateVerifyJWTRequest(req *auth_pb.VerifyTokenRequest) error {
	if req == nil {
		return errRequestNil
	}

	if req.Token == "" {
		return errTokenRequired
	}

	return nil
}
