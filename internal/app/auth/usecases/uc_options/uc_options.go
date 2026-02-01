package uc_options

import (
	"github.com/MediStatTech/auth-service/internal/app/auth/contracts"
)

type Options struct {
	Committer contracts.Committer
	Logger    contracts.Logger

	PositionsRepo contracts.PositionsRepo
	StaffsRepo    contracts.StaffsRepo

	JWT contracts.JWT
}
