package staffs

import (
	"context"

	staffs_pb "github.com/MediStatTech/auth-client/pb/go/services/v1"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/deactivate"
	"github.com/MediStatTech/auth-service/pkg/auth"
)

func (h *Handler) DeactivateStaff(
	ctx context.Context,
	req *staffs_pb.DeactivateStaffRequest,
) (*staffs_pb.DeactivateStaffReply, error) {
	if err := validateDeactivateStaffRequest(req); err != nil {
		return nil, err
	}

	if err := h.commands.DeactivateStaff.Execute(ctx, deactivate.Request{
		Auth:    auth.GetAuth(ctx),
		StaffID: req.StaffId,
	}); err != nil {
		return nil, err
	}

	return &staffs_pb.DeactivateStaffReply{}, nil
}

func validateDeactivateStaffRequest(req *staffs_pb.DeactivateStaffRequest) error {
	if req == nil {
		return errRequestNil
	}

	if req.StaffId == "" {
		return errStaffIDRequired
	}

	return nil
}
