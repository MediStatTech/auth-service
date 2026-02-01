package positions

import (
	"context"

	positions_pb "github.com/MediStatTech/auth-client/pb/go/services/v1"
	"github.com/MediStatTech/auth-service/pkg/auth"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/postions/create"
)

func (h *Handler) CreatePosition(
	ctx context.Context,
	req *positions_pb.CreatePositionRequest,
) (*positions_pb.CreatePositionReply, error) {
	if err := validateCreatePositionRequest(req); err != nil {
		return nil, err
	}

	if err := h.commands.CreatePosition.Execute(ctx, create.Request{
		Auth: auth.GetAuth(ctx),
		Name: req.Name,
	}); err != nil {
		return nil, err
	}

	return &positions_pb.CreatePositionReply{}, nil
}

func validateCreatePositionRequest(req *positions_pb.CreatePositionRequest) error {
	if req == nil {
		return errRequestNil
	}

	if req.Name == "" {
		return errNameRequired
	}

	return nil
}