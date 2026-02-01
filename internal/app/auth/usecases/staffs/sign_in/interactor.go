package sign_in

import (
	"context"
	"errors"

	"github.com/MediStatTech/auth-service/internal/app/auth/contracts"
	"github.com/MediStatTech/auth-service/internal/app/auth/domain"
	"github.com/MediStatTech/auth-service/pkg/password"
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

func (it *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	staffProp, err := it.staffsRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	staff := domain.ReconstituteStaff(staffProp)

	if staff.Status() != domain.StaffStatusActive {
		return nil, errStaffInactive
	}

	if err := comparePassword(req.Password, staff.PasswordHash(), staff.PasswordSalt()); err != nil {
		return nil, err
	}

	token, err := it.jwt.Generate(staff.StaffID(), staff.PositionID())
	if err != nil {
		return nil, errFailedToGenerateJWT
	}

	return &Response{
		Token: token,
	}, nil
}

func comparePassword(
	reqPassword string,
	staffPasswordHash,
	staffPasswordSalt []byte,
) error {
	argon2 := password.NewArgon2idHash()
	if err := argon2.Compare(staffPasswordHash, staffPasswordSalt, []byte(reqPassword)); err != nil {
		if errors.Is(err, password.DoesntMatchError) {
			return errInvalidCredentials
		}
	}

	return nil
}
