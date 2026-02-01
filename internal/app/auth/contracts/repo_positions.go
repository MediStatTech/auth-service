package contracts

import (
	"context"

	"github.com/MediStatTech/auth-service/internal/app/auth/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
)

type PositionsRepo interface {
	ExistsByName(ctx context.Context, name string) (bool, error)
	CreateMut(position *domain.Position) *postgres.Mutation
}
