package create

import (
	"context"
	"time"

	"github.com/MediStatTech/auth-service/internal/app/auth/contracts"
	"github.com/MediStatTech/auth-service/internal/app/auth/domain"
	"github.com/MediStatTech/auth-service/pkg/commitplan"
)

type Interactor struct {
	positionsRepo contracts.PositionsRepo
	committer     contracts.Committer
	logger        contracts.Logger
}

func New(
	positionsRepo contracts.PositionsRepo,
	committer contracts.Committer,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		positionsRepo: positionsRepo,
		committer:     committer,
		logger:        logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) error {
	exists, err := it.positionsRepo.ExistsByName(ctx, req.Name)
	if err != nil {
		return err
	}

	if exists {
		return errPositionAlreadyExists
	}

	position := domain.NewPosition(
		req.Name,
		time.Now().UTC(),
	)

	plan := commitplan.NewPlan()
	plan.AddMut(it.positionsRepo.CreateMut(position))

	if err := it.committer.Apply(ctx, plan); err != nil {
		return errFailedToCreatePosition.SetInternal(err)
	}

	return nil
}
