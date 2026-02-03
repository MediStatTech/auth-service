package repo

import (
	"context"
	"database/sql"

	"github.com/MediStatTech/auth-service/internal/app/auth/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
)

type StaffsRepository struct {
	queries *Queries
}

func NewStaffsRepository(db *sql.DB) *StaffsRepository {
	return &StaffsRepository{
		queries: New(db),
	}
}

// Query operations - use sqlc directly

func (r *StaffsRepository) FindByEmail(ctx context.Context, email string) (domain.StaffProps, error) {
	staff, err := r.queries.GetStaffByEmail(ctx, email)
	if err != nil {
		return domain.StaffProps{}, err
	}

	return toStaffProps(staff), nil
}

func (r *StaffsRepository) FindByID(ctx context.Context, staffID string) (domain.StaffProps, error) {
	staff, err := r.queries.GetStaff(ctx, staffID)
	if err != nil {
		return domain.StaffProps{}, err
	}

	return toStaffProps(staff), nil
}

func (r *StaffsRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	return r.queries.ExistsStaffByEmail(ctx, email)
}

func (r *StaffsRepository) List(ctx context.Context) ([]domain.StaffProps, error) {
	staffs, err := r.queries.ListStaffs(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]domain.StaffProps, 0, len(staffs))
	for _, staff := range staffs {
		result = append(result, toStaffProps(staff))
	}

	return result, nil
}

func (r *StaffsRepository) ListByPosition(ctx context.Context, positionID string) ([]domain.StaffProps, error) {
	staffs, err := r.queries.ListStaffsByPosition(ctx, positionID)
	if err != nil {
		return nil, err
	}

	result := make([]domain.StaffProps, 0, len(staffs))
	for _, staff := range staffs {
		result = append(result, toStaffProps(staff))
	}

	return result, nil
}

func (r *StaffsRepository) Count(ctx context.Context) (int64, error) {
	return r.queries.CountStaffs(ctx)
}

func (r *StaffsRepository) CountByStatus(ctx context.Context, status string) (int64, error) {
	return r.queries.CountStaffsByStatus(ctx, status)
}

// Command operations - return mutations

func (r *StaffsRepository) CreateMut(staff *domain.Staff) *postgres.Mutation {
	return postgres.NewMutation(
		CreateStaff,
		staffToCreateParams(staff)...,
	)
}

func (r *StaffsRepository) UpdateMut(staff *domain.Staff) *postgres.Mutation {
	return postgres.NewMutation(
		UpdateStaff,
		staffToUpdateParams(staff)...,
	)
}

func (r *StaffsRepository) UpdateStatusMut(staffID string, status domain.StaffStatus, updatedAt sql.NullTime) *postgres.Mutation {
	return postgres.NewMutation(
		UpdateStaffStatus,
		staffToUpdateStatusParams(staffID, status, updatedAt)...,
	)
}

func (r *StaffsRepository) DeleteMut(staffID string) *postgres.Mutation {
	return postgres.NewMutation(
		DeleteStaff,
		staffID,
	)
}
