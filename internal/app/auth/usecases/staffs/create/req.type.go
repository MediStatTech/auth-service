package create

import "github.com/MediStatTech/auth-service/pkg/auth"

type Request struct {
	Auth auth.Auth
	Staff Staff
}

type Staff struct {
	PositionID string
	FirstName  string
	LastName   string
	Email      string
	Password   string
}
