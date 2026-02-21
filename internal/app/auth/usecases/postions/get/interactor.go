package get

import (
	"context"

	"github.com/MediStatTech/auth-service/internal/app/auth/contracts"
)

type Interactor struct {
	positionsRepo contracts.PositionsRepo
	logger        contracts.Logger
}

func New(
	positionsRepo contracts.PositionsRepo,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		positionsRepo: positionsRepo,
		logger:        logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (Response, error) {
	positionProps, err := it.positionsRepo.Get(ctx)
	if err != nil {
		return Response{}, errFailedToListPositions.SetInternal(err)
	}

	return Response{
		Positions: positionProps,
	}, nil
}
