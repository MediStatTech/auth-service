package usecases

import (
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/jwt/verify"
	positions_create "github.com/MediStatTech/auth-service/internal/app/auth/usecases/postions/create"
	positions_get "github.com/MediStatTech/auth-service/internal/app/auth/usecases/postions/get"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/create"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/deactivate"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/get"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/retrieve"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/sign_in"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/uc_options"
)

type Facade struct {
	CreatePosition  *positions_create.Interactor
	CreateStaff     *create.Interactor
	DeactivateStaff *deactivate.Interactor
	SignIn          *sign_in.Interactor
	VerifyJWT       *verify.Interactor
	GetPositions    *positions_get.Interactor
	GetStaffs       *get.Interactor
	RetrieveStaff   *retrieve.Interactor
}

func New(o *uc_options.Options) *Facade {
	return &Facade{
		CreatePosition:  positions_create.New(o.PositionsRepo, o.Committer, o.Logger),
		CreateStaff:     create.New(o.StaffsRepo, o.Committer, o.Logger),
		DeactivateStaff: deactivate.New(o.StaffsRepo, o.Committer, o.Logger),
		SignIn:          sign_in.New(o.JWT, o.StaffsRepo, o.Committer, o.Logger),
		VerifyJWT:       verify.New(o.JWT, o.Logger),
		GetPositions:    positions_get.New(o.PositionsRepo, o.Logger),
		GetStaffs:       get.New(o.StaffsRepo, o.Logger),
		RetrieveStaff:   retrieve.New(o.StaffsRepo, o.Logger),
	}
}
