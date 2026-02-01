package pkg

import (
	"context"

	"github.com/MediStatTech/commitplan"
	"github.com/MediStatTech/jwt"
	"github.com/MediStatTech/logger"
	"github.com/MediStatTech/auth-service/pkg/config"
)

type Facade struct {
	Committer *commitplan.Facade
	JWT       *jwt.JWT
	Logger    *logger.Logger
	Config    *config.Config
}


func New(ctx context.Context) (*Facade, error) {
	return &Facade{}, nil
}