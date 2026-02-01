package usecases

import (
	positions_create "github.com/MediStatTech/auth-service/internal/app/auth/usecases/postions/create"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/create"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/deactivate"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/sign_in"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/uc_options"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/jwt/verify"
)

type Facade struct {
	CreatePosition  *positions_create.Interactor
	CreateStaff     *create.Interactor
	DeactivateStaff *deactivate.Interactor
	SignIn          *sign_in.Interactor
	VerifyJWT       *verify.Interactor
}

func New(o *uc_options.Options) *Facade {
	return &Facade{
		CreatePosition:  positions_create.New(o.PositionsRepo, o.Committer, o.Logger),
		CreateStaff:     create.New(o.StaffsRepo, o.Committer, o.Logger),
		DeactivateStaff: deactivate.New(o.StaffsRepo, o.Committer, o.Logger),
		SignIn:          sign_in.New(o.JWT, o.StaffsRepo, o.Committer, o.Logger),
		VerifyJWT:       verify.New(o.JWT, o.Logger),
	}
}
