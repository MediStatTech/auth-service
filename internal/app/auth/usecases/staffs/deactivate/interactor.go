package deactivate

import (
	"context"

	"github.com/MediStatTech/auth-service/internal/app/auth/contracts"
	"github.com/MediStatTech/auth-service/internal/app/auth/domain"
	"github.com/MediStatTech/auth-service/pkg/commitplan"
)

type Interactor struct {
	staffsRepo contracts.StaffsRepo
	committer  contracts.Committer
	logger     contracts.Logger
}

func New(
	staffsRepo contracts.StaffsRepo,
	committer contracts.Committer,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		staffsRepo: staffsRepo,
		committer:  committer,
		logger:     logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) error {
	plan := commitplan.NewPlan()

	//todo: validate permissions
	staffProp, err := it.staffsRepo.FindByID(ctx, req.StaffID)
	if err != nil {
		return errFailedToFindStaff
	}

	staff := domain.ReconstituteStaff(staffProp)

	staff.SetInactiveStatus()
	plan.AddMut(it.staffsRepo.UpdateMut(staff))

	if err := it.committer.Apply(ctx, plan); err != nil {
		return errFailedToUpdateStaff.SetInternal(err)
	}

	return nil
}
