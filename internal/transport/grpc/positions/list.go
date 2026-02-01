package positions

import (
	"context"

	positions_pb "github.com/MediStatTech/auth-client/pb/go/services/v1"
)

func (h *Handler) ListPositions(
	ctx context.Context,
	req *positions_pb.ListPositionsRequest,
) (*positions_pb.ListPositionsReply, error) {
	return nil, nil
}