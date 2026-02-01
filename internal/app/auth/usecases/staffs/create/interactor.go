package create

import (
	"context"
	"time"

	"github.com/MediStatTech/auth-service/internal/app/auth/contracts"
	"github.com/MediStatTech/auth-service/internal/app/auth/domain"
	"github.com/MediStatTech/auth-service/pkg/commitplan"
	"github.com/MediStatTech/auth-service/pkg/password"
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
	exists, err := it.staffsRepo.ExistsByEmail(ctx, req.Staff.Email)
	if err != nil {
		return err
	}

	if exists {
		return errEmailAlreadyExists
	}

	//TODO: Validate permissions

	hashSalt, err := hash(req.Staff.Password)
	if err != nil {
		return errFailedToHashPassword
	}

	staff := domain.NewStaff(
		req.Staff.PositionID,
		req.Staff.FirstName,
		req.Staff.LastName,
		req.Staff.Email,
		hashSalt.Hash,
		hashSalt.Salt,
		time.Now().UTC(),
	)

	plan := commitplan.NewPlan()
	plan.AddMut(it.staffsRepo.CreateMut(staff))

	if err := it.committer.Apply(ctx, plan); err != nil {
		return errFailedToCreateStaff.SetInternal(err)
	}

	return nil
}

func hash(staffPassword string) (*password.HashSalt, error) {
	argon2 := password.NewArgon2idHash()
	hashSalt, err := argon2.GenerateHash([]byte(staffPassword), nil)

	return hashSalt, err
}
