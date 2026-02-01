package jwt

import (
	s_options "github.com/MediStatTech/auth-service/internal/app/options"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/jwt/verify"
	"github.com/MediStatTech/logger"
)

type Handler struct {
	log      *logger.Logger
	commands *Commands
}

type Commands struct {
	//JWT
	Verify *verify.Interactor
}

func New(opts *s_options.Options) *Handler {
	return &Handler{
		log: opts.PKG.Logger,
		commands: &Commands{
			Verify: opts.App.Auth.VerifyJWT,
		},
	}
}
