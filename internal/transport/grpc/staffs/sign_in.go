package staffs

import (
	"context"

	staffs_pb "github.com/MediStatTech/auth-client/pb/go/services/v1"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/sign_in"
)

func (h *Handler) SignIn(
	ctx context.Context,
	req *staffs_pb.SignInRequest,
) (*staffs_pb.SignInReply, error) {

	if err := validateSignInRequest(req); err != nil {
		return nil, err
	}

	reply, err := h.commands.SignIn.Execute(ctx, sign_in.Request{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &staffs_pb.SignInReply{
		Token: reply.Token,
	}, nil
}

func validateSignInRequest(req *staffs_pb.SignInRequest) error {
	if req == nil {
		return errRequestNil
	}

	if req.Email == "" {
		return errEmailRequired
	}

	return nil
}
