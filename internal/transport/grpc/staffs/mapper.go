package staffs

import (
	pb_models "github.com/MediStatTech/auth-client/pb/go/models/v1"
	"github.com/MediStatTech/auth-service/internal/app/auth/domain"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/create"
)

func pbToCreateStaffModel(req *pb_models.Staff_Create) create.Staff {
	return create.Staff{
		PositionID: req.PositionId,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Email:      req.Email,
		Password:   req.Password,
	}
}

func staffPropsToPb(staff *domain.StaffProps) *pb_models.Staff_Read {
	return &pb_models.Staff_Read{
		StaffId:        staff.StaffID,
		PositionId:     staff.PositionID,
		FirstName:      staff.FirstName,
		LastName:       staff.LastName,
		SelfieUrl:      staff.SelfieURL,
		SelfieThumbUrl: staff.SelfieThumbURL,
		Status:         staff.Status,
		Email:          staff.Email,
	}
}
