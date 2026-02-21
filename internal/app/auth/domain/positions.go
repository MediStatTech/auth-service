package domain

import (
	"time"

	"github.com/google/uuid"
)

type PositionProps struct {
	PositionID string
	Name       string
	UpdatedAt  *time.Time
	CreatedAt  time.Time
}

type Position struct {
	positionID string
	name       string
	updatedAt  *time.Time
	createdAt  time.Time
}

func NewPosition(
	name string,
	createdAt time.Time,
) *Position {
	return &Position{
		positionID: uuid.NewString(),
		name:       name,
		createdAt:  createdAt,
	}
}

func ReconstitutePosition(p PositionProps) *Position {
	return &Position{
		positionID: p.PositionID,
		name:       p.Name,
		createdAt:  p.CreatedAt,
		updatedAt:  p.UpdatedAt,
	}
}

func (p *Position) PositionID() string   { return p.positionID }
func (p *Position) Name() string         { return p.name }
func (p *Position) UpdatedAt() *time.Time { return p.updatedAt }
func (p *Position) CreatedAt() time.Time { return p.createdAt }

func (p *Position) SetName(name string)  {
	p.name = name
}

func (p *Position) SetUpdatedAt(updatedAt time.Time) {
	p.updatedAt = &updatedAt
}
