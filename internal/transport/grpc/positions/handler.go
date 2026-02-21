package positions

import (
	positions_create "github.com/MediStatTech/auth-service/internal/app/auth/usecases/postions/create"
	positions_get "github.com/MediStatTech/auth-service/internal/app/auth/usecases/postions/get"
	s_options "github.com/MediStatTech/auth-service/internal/app/options"
	"github.com/MediStatTech/logger"
)

type Handler struct {
	log      *logger.Logger
	commands *Commands
	//queries
}

type Commands struct {
	//Positions
	CreatePosition *positions_create.Interactor
	GetPositions   *positions_get.Interactor
}

func New(opts *s_options.Options) *Handler {
	return &Handler{
		log: opts.PKG.Logger,
		commands: &Commands{
			CreatePosition: opts.App.Auth.CreatePosition,
			GetPositions:   opts.App.Auth.GetPositions,
		},
	}
}
