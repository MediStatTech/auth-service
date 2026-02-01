package app

import (
	"fmt"

	"github.com/MediStatTech/auth-service/internal/app/auth"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases"
	"github.com/MediStatTech/auth-service/pkg"
)

type Facade struct {
	Auth *usecases.Facade
}

func New(pkg *pkg.Facade) (*Facade, error) {
	auth, err := auth.New(pkg)
	if err != nil {
		return nil, fmt.Errorf("failed to create auth: %w", err)
	}

	return &Facade{
		Auth: auth.UseCases,
	}, nil
}
