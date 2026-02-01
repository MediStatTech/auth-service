package staffs

import (
	staffs_pb "github.com/MediStatTech/auth-client/pb/go/services/v1"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/create"
)

func formatCreateStaffModel(req *staffs_pb.CreateStaffRequest) create.Staff {
	return create.Staff{
		PositionID: req.PositionId,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Email:      req.Email,
		Password:   req.Password,
	}
}
