package domain

import (
	"time"

	"github.com/google/uuid"
)

type StaffStatus string

const (
	StaffStatusActive    StaffStatus = "active"
	StaffStatusInactive  StaffStatus = "inactive"
	StaffStatusSuspended StaffStatus = "suspended"
)

func (s StaffStatus) String() string {
	return string(s)
}

type StaffProps struct {
	StaffID        string
	PositionID     string
	FirstName      string
	LastName       string
	email          string
	PasswordHash   []byte
	PasswordSalt   []byte
	Status         string
	SelfieURL      *string
	SelfieThumbURL *string
	UpdatedAt      *time.Time
	CreatedAt      time.Time
}

type Staff struct {
	staffID        string
	positionID     string
	firstName      string
	lastName       string
	email          string
	passwordHash   []byte
	passwordSalt   []byte
	status         string
	selfieURL      *string
	selfieThumbURL *string
	updatedAt      *time.Time
	createdAt      time.Time
}

func NewStaff(
	positionID string,
	firstName, lastName, email string,
	passwordHash, passwordSalt []byte,
	createdAt time.Time,
) *Staff {
	return &Staff{
		staffID:      uuid.NewString(),
		positionID:   positionID,
		firstName:    firstName,
		lastName:     lastName,
		email:        email,
		passwordHash: passwordHash,
		passwordSalt: passwordSalt,
		status:       StaffStatusActive.String(),
		createdAt:    createdAt,
	}
}

func ReconstituteStaff(p StaffProps) *Staff {
	return &Staff{
		staffID:        p.StaffID,
		positionID:     p.PositionID,
		firstName:      p.FirstName,
		lastName:       p.LastName,
		email:          p.email,
		passwordHash:   p.PasswordHash,
		passwordSalt:   p.PasswordSalt,
		status:         p.Status,
		selfieURL:      p.SelfieURL,
		selfieThumbURL: p.SelfieThumbURL,
		createdAt:      p.CreatedAt,
		updatedAt:      p.UpdatedAt,
	}
}

func (s *Staff) StaffID() string         { return s.staffID }
func (s *Staff) PositionID() string      { return s.positionID }
func (s *Staff) FirstName() string       { return s.firstName }
func (s *Staff) LastName() string        { return s.lastName }
func (s *Staff) Email() string           { return s.email }
func (s *Staff) PasswordHash() []byte    { return s.passwordHash }
func (s *Staff) PasswordSalt() []byte    { return s.passwordSalt }
func (s *Staff) Status() StaffStatus     { return StaffStatus(s.status) }
func (s *Staff) SelfieURL() *string      { return s.selfieURL }
func (s *Staff) SelfieThumbURL() *string { return s.selfieThumbURL }
func (s *Staff) UpdatedAt() *time.Time   { return s.updatedAt }
func (s *Staff) CreatedAt() time.Time    { return s.createdAt }

func (s *Staff) SetFirstName(firstName string) *Staff {
	s.firstName = firstName
	return s
}

func (s *Staff) SetLastName(lastName string) *Staff {
	s.lastName = lastName
	return s
}

func (s *Staff) SetPasswordHash(passwordHash []byte) *Staff {
	s.passwordHash = passwordHash
	return s
}

func (s *Staff) SetPasswordSalt(passwordSalt []byte) *Staff {
	s.passwordSalt = passwordSalt
	return s
}

func (s *Staff) SetStatus(status StaffStatus) *Staff {
	s.status = status.String()
	return s
}

func (s *Staff) SetActiveStatus() *Staff {
	s.status = StaffStatusActive.String()
	return s
}

func (s *Staff) SetInactiveStatus() *Staff {
	s.status = StaffStatusInactive.String()
	return s
}

func (s *Staff) SetSuspendedStatus() *Staff {
	s.status = StaffStatusSuspended.String()
	return s
}

func (s *Staff) SetSelfieURL(selfieURL string) *Staff {
	s.selfieURL = &selfieURL
	return s
}

func (s *Staff) SetSelfieThumbURL(selfieThumbURL string) *Staff {
	s.selfieThumbURL = &selfieThumbURL
	return s
}

func (s *Staff) SetPositionID(positionID string) *Staff {
	s.positionID = positionID
	return s
}

func (s *Staff) SetUpdatedAt(updatedAt time.Time) *Staff {
	s.updatedAt = &updatedAt
	return s
}

func (s *Staff) FullName() string {
	return s.firstName + " " + s.lastName
}
