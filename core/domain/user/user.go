package user

import (
	"errors"
	"time"
)

// Role represents the user's role in the fitworks system.
type Role string

const (
	RoleTrainer Role = "trainer"
	RoleTrainee Role = "trainee"
	RoleNone    Role = "none"
)

// Sentinel domain errors.
var (
	ErrNotFound      = errors.New("user not found")
	ErrEmailConflict = errors.New("email already in use")
	ErrInvalidRole   = errors.New("invalid role: must be trainer, trainee, or none")
)

// IsValidRole reports whether r is a recognised Role value.
func IsValidRole(r Role) bool {
	switch r {
	case RoleTrainer, RoleTrainee, RoleNone:
		return true
	}
	return false
}

// User is the core domain entity.
type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
}
