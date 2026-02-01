package staffs

import (
	"context"

	staffs_pb "github.com/MediStatTech/auth-client/pb/go/services/v1"
)

func (h *Handler) FindStaff(
	ctx context.Context,
	req *staffs_pb.FindStaffRequest,
) (*staffs_pb.FindStaffReply, error) {
	return nil, nil
}