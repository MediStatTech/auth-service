package staffs

import (
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/create"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/deactivate"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/sign_in"
	s_options "github.com/MediStatTech/auth-service/internal/app/options"
	"github.com/MediStatTech/logger"
)

type Handler struct {
	log      *logger.Logger
	commands *Commands
	//queries
}

type Commands struct {
	//Staffs
	CreateStaff     *create.Interactor
	DeactivateStaff *deactivate.Interactor
	SignIn          *sign_in.Interactor
	// FindStaff       *find.Interactor
	// ListStaff       *list.Interactor
}

func New(opts *s_options.Options) *Handler {
	return &Handler{
		log: opts.PKG.Logger,
		commands: &Commands{
			CreateStaff:     opts.App.Auth.CreateStaff,
			DeactivateStaff: opts.App.Auth.DeactivateStaff,
			SignIn:          opts.App.Auth.SignIn,
		},
	}
}
