package positions

import (
	pb_models "github.com/MediStatTech/auth-client/pb/go/models/v1"
	"github.com/MediStatTech/auth-service/internal/app/auth/domain"
)

func positionPropsToPb(positions []domain.PositionProps) []*pb_models.Position_Read {
	pbPositions := make([]*pb_models.Position_Read, len(positions))
	for i, position := range positions {
		pbPositions[i] = &pb_models.Position_Read{
			PositionId: position.PositionID,
			Name:       position.Name,
		}
	}
	return pbPositions
}
