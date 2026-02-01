package auth

import (
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/uc_options"
	"github.com/MediStatTech/auth-service/pkg"
)

type Facade struct {
	pkg      *pkg.Facade
	UseCases *usecases.Facade
}

func New(pkg *pkg.Facade) (*Facade, error) {

	useCasesInstance := usecases.New(&uc_options.Options{
		Committer:     pkg.Committer,
		Logger:        pkg.Logger,
		// PositionsRepo: pkg.PositionsRepo,
		// StaffsRepo:    pkg.StaffsRepo, // TODO: Implement
		JWT:           pkg.JWT,
	})

	return &Facade{
		pkg:      pkg,
		UseCases: useCasesInstance,
	}, nil
}
