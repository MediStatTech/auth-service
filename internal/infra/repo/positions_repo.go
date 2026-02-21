package repo

import (
	"context"
	"database/sql"

	"github.com/MediStatTech/auth-service/internal/app/auth/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
)

type PositionsRepository struct {
	queries *Queries
}

func NewPositionsRepository(db *sql.DB) *PositionsRepository {
	return &PositionsRepository{
		queries: New(db),
	}
}

func (r *PositionsRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	return r.queries.ExistsPositionByName(ctx, name)
}

func (r *PositionsRepository) GetByID(ctx context.Context, positionID string) (domain.PositionProps, error) {
	pos, err := r.queries.GetPosition(ctx, positionID)
	if err != nil {
		return domain.PositionProps{}, err
	}

	return toPositionProps(pos), nil
}

func (r *PositionsRepository) Get(ctx context.Context) ([]domain.PositionProps, error) {
	positions, err := r.queries.ListPositions(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.PositionProps, 0, len(positions))
	for _, pos := range positions {
		result = append(result, toPositionProps(pos))
	}

	return result, nil
}

func (r *PositionsRepository) GetByName(ctx context.Context, name string) (domain.PositionProps, error) {
	pos, err := r.queries.GetPositionByName(ctx, name)
	if err != nil {
		return domain.PositionProps{}, err
	}

	return toPositionProps(pos), nil
}

func (r *PositionsRepository) List(ctx context.Context) ([]domain.PositionProps, error) {
	positions, err := r.queries.ListPositions(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.PositionProps, 0, len(positions))
	for _, pos := range positions {
		result = append(result, toPositionProps(pos))
	}

	return result, nil
}

func (r *PositionsRepository) Count(ctx context.Context) (int64, error) {
	return r.queries.CountPositions(ctx)
}

// Command operations - return mutations

func (r *PositionsRepository) CreateMut(position *domain.Position) *postgres.Mutation {
	return postgres.NewMutation(
		CreatePosition,
		positionToCreateParams(position)...,
	)
}

func (r *PositionsRepository) UpdateMut(position *domain.Position) *postgres.Mutation {
	return postgres.NewMutation(
		UpdatePosition,
		positionToUpdateParams(position)...,
	)
}

func (r *PositionsRepository) DeleteMut(positionID string) *postgres.Mutation {
	return postgres.NewMutation(
		DeletePosition,
		positionID,
	)
}
