package retrieve

import (
	"context"

	"github.com/MediStatTech/auth-service/internal/app/auth/contracts"
)

type Interactor struct {
	staffsRepo contracts.StaffsRepo
	logger     contracts.Logger
}

func New(
	staffsRepo contracts.StaffsRepo,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		staffsRepo: staffsRepo,
		logger:     logger,
	}
}

func (it *Interactor) Execute(ctx context.Context, req Request) (Response, error) {
	staffProp, err := it.staffsRepo.Find(ctx, req.StaffID)
	if err != nil {
		return Response{}, errFailedToFindStaff.SetInternal(err)
	}

	return Response{
		Staff: staffProp,
	}, nil
}
