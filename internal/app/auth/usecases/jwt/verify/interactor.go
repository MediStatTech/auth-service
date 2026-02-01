package verify

import (
	"context"

	"github.com/MediStatTech/auth-service/internal/app/auth/contracts"
)

type Interactor struct {
	jwt contracts.JWT
	logger contracts.Logger
}

func New(
	jwt contracts.JWT,
	logger contracts.Logger,
) *Interactor {
	return &Interactor{
		jwt: jwt,
		logger: logger,
	}
}

func (i *Interactor) Execute(ctx context.Context, req Request) (*Response, error) {
	claims, err := i.jwt.Verify(req.Token)
	if err != nil {
		return nil, errFailedToVerifyJWT.SetInternal(err)
	}

	return &Response{
		StaffID:    claims.StaffID,
		PositionID: claims.PositionID,
	}, nil
}