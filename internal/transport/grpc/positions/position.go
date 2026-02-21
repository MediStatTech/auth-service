package positions

import (
	"context"

	pb_services "github.com/MediStatTech/auth-client/pb/go/services/v1"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/postions/create"
	positions_get "github.com/MediStatTech/auth-service/internal/app/auth/usecases/postions/get"
)

func (h *Handler) PositionCreate(
	ctx context.Context,
	req *pb_services.PositionCreateRequest,
) (*pb_services.PositionCreateReply, error) {
	if req == nil {
		return nil, errRequestNil
	}

	if err := h.commands.CreatePosition.Execute(ctx, create.Request{
		Name: req.Postion.Name,
	}); err != nil {
		return nil, err
	}

	return &pb_services.PositionCreateReply{}, nil
}

func (h *Handler) PositionGet(
	ctx context.Context,
	req *pb_services.PositionGetRequest,
) (*pb_services.PositionGetReply, error) {
	reply, err := h.commands.GetPositions.Execute(ctx, positions_get.Request{})
	if err != nil {
		return nil, err
	}

	return &pb_services.PositionGetReply{
		Positions: positionPropsToPb(reply.Positions),
	}, nil
}
