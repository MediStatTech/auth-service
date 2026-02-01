package staffs

import (
	"context"

	staffs_pb "github.com/MediStatTech/auth-client/pb/go/services/v1"
)

func (h *Handler) ListStaffs(
	ctx context.Context,
	req *staffs_pb.ListStaffsRequest,
) (*staffs_pb.ListStaffsReply, error) {
	return nil, nil
}