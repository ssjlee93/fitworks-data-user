package user

import "context"

// Repository is the driven port — implemented by pkg/postgres.
type Repository interface {
	FindByID(ctx context.Context, id int64) (*User, error)
	FindAll(ctx context.Context) ([]*User, error)
	Save(ctx context.Context, u *User) (*User, error)
	Update(ctx context.Context, u *User) (*User, error)
	Delete(ctx context.Context, id int64) error
}

// Service is the driving port — implemented by core/service/user.
type Service interface {
	GetUser(ctx context.Context, id int64) (*User, error)
	ListUsers(ctx context.Context) ([]*User, error)
	CreateUser(ctx context.Context, cmd CreateCommand) (*User, error)
	UpdateUser(ctx context.Context, id int64, cmd UpdateCommand) (*User, error)
	DeleteUser(ctx context.Context, id int64) error
}

// CreateCommand carries validated input for user creation.
type CreateCommand struct {
	FirstName string
	LastName  string
	Email     string
	Role      Role
}

// UpdateCommand carries validated input for user updates.
// Pointer fields allow partial updates — nil means "no change".
type UpdateCommand struct {
	FirstName *string
	LastName  *string
	Email     *string
	Role      *Role
}
