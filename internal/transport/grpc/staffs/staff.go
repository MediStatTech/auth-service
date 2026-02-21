package staffs

import (
	"context"

	pb_services "github.com/MediStatTech/auth-client/pb/go/services/v1"
	pb_models "github.com/MediStatTech/auth-client/pb/go/models/v1"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/create"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/get"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/deactivate"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/retrieve"
	"github.com/MediStatTech/auth-service/internal/app/auth/usecases/staffs/sign_in"
)

func (h *Handler) StaffCreate(
	ctx context.Context,
	req *pb_services.StaffCreateRequest,
) (*pb_services.StaffCreateReply, error) {
	if err := h.commands.CreateStaff.Execute(ctx, create.Request{
		Staff: pbToCreateStaffModel(req.Staff),
	}); err != nil {
		return nil, err
	}

	return &pb_services.StaffCreateReply{}, nil
}

func (h *Handler) StaffDeactivate(
	ctx context.Context,
	req *pb_services.StaffDeactivateRequest,
) (*pb_services.StaffDeactivateReply, error) {
	if err := h.commands.DeactivateStaff.Execute(ctx, deactivate.Request{
		StaffID: req.StaffId,
	}); err != nil {
		return nil, err
	}
	return &pb_services.StaffDeactivateReply{}, nil
}

func (h *Handler) StaffRetrieve(
	ctx context.Context,
	req *pb_services.StaffRetrieveRequest,
) (*pb_services.StaffRetrieveReply, error) {
	reply, err := h.commands.RetrieveStaff.Execute(ctx, retrieve.Request{
		StaffID: req.StaffId,
	})
	if err != nil {
		return nil, err
	}

	return &pb_services.StaffRetrieveReply{
		Staff: staffPropsToPb(&reply.Staff),
	}, nil
}

func (h *Handler) StaffGet(
	ctx context.Context,
	req *pb_services.StaffGetRequest,
) (*pb_services.StaffGetReply, error) {
	reply, err := h.commands.GetStaffs.Execute(ctx, get.Request{})
	if err != nil {
		return nil, err
	}

	pbStaffs := make([]*pb_models.Staff_Read, len(reply.Staffs))
	for i, staff := range reply.Staffs {
		pbStaffs[i] = staffPropsToPb(&staff)
	}

	return &pb_services.StaffGetReply{
		Staffs: pbStaffs,
	}, nil
}

func (h *Handler) SignIn(
	ctx context.Context,
	req *pb_services.SignInRequest,
) (*pb_services.SignInReply, error) {
	reply, err := h.commands.SignIn.Execute(ctx, sign_in.Request{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &pb_services.SignInReply{
		Token: reply.Token,
	}, nil
}
