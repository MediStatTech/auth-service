package deactivate

import "github.com/MediStatTech/auth-service/pkg/auth"

type Request struct {
	Auth auth.Auth
	StaffID string
}