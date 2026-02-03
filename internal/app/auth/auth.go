package auth

import (
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/uc_options"
	"github.com/MediStatTech/auth-service/internal/infra/repo"
	"github.com/MediStatTech/auth-service/pkg"
)

type Facade struct {
	pkg      *pkg.Facade
	UseCases *usecases.Facade
}

func New(pkg *pkg.Facade) (*Facade, error) {
	positionsRepo := repo.NewPositionsRepository(pkg.Postgres.DB)
	staffsRepo := repo.NewStaffsRepository(pkg.Postgres.DB)

	useCasesInstance := usecases.New(&uc_options.Options{
		Committer:     pkg.Committer,
		Logger:        pkg.Logger,
		PositionsRepo: positionsRepo,
		StaffsRepo:    staffsRepo,
		JWT:           pkg.JWT,
	})

	return &Facade{
		pkg:      pkg,
		UseCases: useCasesInstance,
	}, nil
}
