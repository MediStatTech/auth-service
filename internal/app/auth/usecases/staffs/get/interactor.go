package get

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
	staffProps, err := it.staffsRepo.List(ctx)
	if err != nil {
		return Response{}, errFailedToListStaffs.SetInternal(err)
	}

	return Response{
		Staffs: staffProps,
	}, nil
}
