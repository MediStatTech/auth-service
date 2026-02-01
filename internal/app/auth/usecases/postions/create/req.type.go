package create

import "github.com/MediStatTech/auth-service/pkg/auth"

type Request struct {
	Auth auth.Auth
	Name string
}
