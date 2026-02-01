package staffs

import (
	"context"

	staffs_pb "github.com/MediStatTech/auth-client/pb/go/services/v1"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/create"
	"github.com/MediStatTech/auth-service/pkg/auth"
)

func (h *Handler) CreateStaff(
	ctx context.Context,
	req *staffs_pb.CreateStaffRequest,
) (*staffs_pb.CreateStaffReply, error) {
	if err := validateCreateStaffRequest(req); err != nil {
		return nil, err
	}

	if err := h.commands.CreateStaff.Execute(ctx, create.Request{
		Staff: formatCreateStaffModel(req),
		Auth: auth.GetAuth(ctx),
	});err != nil {
		return nil, err
	}

	return &staffs_pb.CreateStaffReply{}, nil
}

func validateCreateStaffRequest(req *staffs_pb.CreateStaffRequest) error {
	if req == nil {
		return errRequestNil
	}

	if req.Email == "" {
		return errEmailRequired
	}

	if req.Password == "" {
		return errPasswordRequired
	}

	if req.PositionId == "" {
		return errPositionIDRequired
	}

	if req.FirstName == "" {
		return errFirstNameRequired
	}

	if req.LastName == "" {
		return errLastNameRequired
	}

	return nil
}