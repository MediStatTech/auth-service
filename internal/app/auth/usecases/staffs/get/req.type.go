package get

import "github.com/MediStatTech/auth-service/internal/app/auth/domain"

type Request struct {
}

type Response struct {
	Staffs []domain.StaffProps
}
