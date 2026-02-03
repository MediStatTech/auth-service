package repo

import (
	"database/sql"

	"github.com/MediStatTech/auth-service/internal/app/auth/domain"
)

// ============================================================================
// Positions Mappers
// ============================================================================

func toPositionProps(pos Position) domain.PositionProps {
	props := domain.PositionProps{
		PositionID: pos.PositionID,
		Name:       pos.Name,
		CreatedAt:  pos.CreatedAt,
	}

	if pos.UpdatedAt.Valid {
		props.UpdatedAt = &pos.UpdatedAt.Time
	}

	return props
}

func positionToCreateParams(position *domain.Position) []any {
	return []any{
		position.PositionID(),
		position.Name(),
		position.CreatedAt(),
	}
}

func positionToUpdateParams(position *domain.Position) []any {
	var updatedAt sql.NullTime
	if position.UpdatedAt() != nil {
		updatedAt = sql.NullTime{
			Time:  *position.UpdatedAt(),
			Valid: true,
		}
	}

	return []any{
		position.PositionID(),
		position.Name(),
		updatedAt,
	}
}

// ============================================================================
// Staffs Mappers
// ============================================================================

func toStaffProps(staff Staff) domain.StaffProps {
	props := domain.StaffProps{
		StaffID:      staff.StaffID,
		PositionID:   staff.PositionID,
		FirstName:    staff.FirstName,
		LastName:     staff.LastName,
		Email:        staff.Email,
		PasswordHash: staff.PasswordHash,
		PasswordSalt: staff.PasswordSalt,
		Status:       staff.Status,
		CreatedAt:    staff.CreatedAt,
	}

	if staff.SelfieUrl.Valid {
		props.SelfieURL = &staff.SelfieUrl.String
	}

	if staff.SelfieThumbUrl.Valid {
		props.SelfieThumbURL = &staff.SelfieThumbUrl.String
	}

	if staff.UpdatedAt.Valid {
		props.UpdatedAt = &staff.UpdatedAt.Time
	}


	return props
}

func staffToCreateParams(staff *domain.Staff) []any {
	var selfieURL, selfieThumbURL sql.NullString

	if staff.SelfieURL() != nil {
		selfieURL = sql.NullString{
			String: *staff.SelfieURL(),
			Valid:  true,
		}
	}

	if staff.SelfieThumbURL() != nil {
		selfieThumbURL = sql.NullString{
			String: *staff.SelfieThumbURL(),
			Valid:  true,
		}
	}

	return []any{
		staff.StaffID(),
		staff.PositionID(),
		staff.FirstName(),
		staff.LastName(),
		staff.Email(),
		staff.PasswordHash(),
		staff.PasswordSalt(),
		staff.Status().String(),
		selfieURL,
		selfieThumbURL,
		staff.CreatedAt(),
	}
}

func staffToUpdateParams(staff *domain.Staff) []any {
	var selfieURL, selfieThumbURL sql.NullString
	var updatedAt sql.NullTime

	if staff.SelfieURL() != nil {
		selfieURL = sql.NullString{
			String: *staff.SelfieURL(),
			Valid:  true,
		}
	}

	if staff.SelfieThumbURL() != nil {
		selfieThumbURL = sql.NullString{
			String: *staff.SelfieThumbURL(),
			Valid:  true,
		}
	}

	if staff.UpdatedAt() != nil {
		updatedAt = sql.NullTime{
			Time:  *staff.UpdatedAt(),
			Valid: true,
		}
	}

	return []any{
		staff.StaffID(),
		staff.FirstName(),
		staff.LastName(),
		staff.Email(),
		staff.PasswordHash(),
		staff.PasswordSalt(),
		staff.Status().String(),
		selfieURL,
		selfieThumbURL,
		updatedAt,
	}
}

func staffToUpdateStatusParams(staffID string, status domain.StaffStatus, updatedAt sql.NullTime) []any {
	return []any{
		staffID,
		status.String(),
		updatedAt,
	}
}
