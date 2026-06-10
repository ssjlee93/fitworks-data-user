package http

import (
	"fmt"

	domain "github.com/ssjlee93/fitworks-data-user/core/domain/user"
)

// CreateUserRequest is the JSON body for POST /api/v1/users.
type CreateUserRequest struct {
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email"`
	Role      domain.Role `json:"role"`
}

// Validate performs basic field-presence checks.
func (r CreateUserRequest) Validate() error {
	switch {
	case r.FirstName == "":
		return fmt.Errorf("field %q is required", "first_name")
	case r.LastName == "":
		return fmt.Errorf("field %q is required", "last_name")
	case r.Email == "":
		return fmt.Errorf("field %q is required", "email")
	case r.Role == "":
		return fmt.Errorf("field %q is required", "role")
	}
	return nil
}

// UpdateUserRequest is the JSON body for PUT/PATCH /api/v1/users/{id}.
// All fields are optional — nil means "leave unchanged".
type UpdateUserRequest struct {
	FirstName *string      `json:"first_name"`
	LastName  *string      `json:"last_name"`
	Email     *string      `json:"email"`
	Role      *domain.Role `json:"role"`
}
