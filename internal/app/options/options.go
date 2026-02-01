package options

import (
	"github.com/MediStatTech/auth-service/internal/app"
	"github.com/MediStatTech/auth-service/pkg"
)

type Options struct {
	App *app.Facade
	PKG *pkg.Facade
}
