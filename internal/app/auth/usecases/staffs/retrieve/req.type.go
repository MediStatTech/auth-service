package retrieve

import "github.com/MediStatTech/auth-service/internal/app/auth/domain"

type Request struct {
	StaffID string
}

type Response struct {
	Staff domain.StaffProps
}
