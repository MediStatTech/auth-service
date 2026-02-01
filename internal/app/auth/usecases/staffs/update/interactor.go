package sign_in

import (
	"github.com/MediStatTech/auth-service/internal/app/auth/contracts"
)

type Interactor struct {
	jwt        contracts.JWT
	staffsRepo contracts.StaffsRepo
	committer  contracts.Committer
	logger     contracts.Logger
}

func New(
	jwt contracts.JWT,
	staffsRepo contracts.StaffsRepo,
	committer contracts.Committer,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		jwt:        jwt,
		staffsRepo: staffsRepo,
		committer:  committer,
		logger:     logger,
	}
}
