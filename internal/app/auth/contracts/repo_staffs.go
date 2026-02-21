package contracts

import (
	"context"

	"github.com/MediStatTech/auth-service/internal/app/auth/domain"
	"github.com/MediStatTech/commitplan/drivers/postgres"
)

type StaffsRepo interface {
	FindByEmail(ctx context.Context, email string) (domain.StaffProps, error)
	FindByID(ctx context.Context, staffID string) (domain.StaffProps, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	List(ctx context.Context) ([]domain.StaffProps, error)
	Find(ctx context.Context, staffID string) (domain.StaffProps, error)
	CreateMut(staff *domain.Staff) *postgres.Mutation
	UpdateMut(staff *domain.Staff) *postgres.Mutation
}
