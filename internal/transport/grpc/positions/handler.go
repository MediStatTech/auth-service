package positions

import (
	positions_create "github.com/MediStatTech/auth-service/internal/app/auth/usecases/postions/create"
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
	// ListPositions   *positions_list.Interactor
	// FindPosition    *positions_find.Interactor
}

func New(opts *s_options.Options) *Handler {
	return &Handler{
		log: opts.PKG.Logger,
		commands: &Commands{
			CreatePosition: opts.App.Auth.CreatePosition,
		},
	}
}
