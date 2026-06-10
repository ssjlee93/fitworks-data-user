package user

import (
	"context"
	"fmt"

	domain "github.com/ssjlee93/fitworks-data-user/core/domain/user"
)

// Service implements domain.Service.
type Service struct {
	repo domain.Repository
}

// New constructs a Service with the given Repository.
func New(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

// GetUser returns a single user by ID.
func (s *Service) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service.GetUser: %w", err)
	}
	return u, nil
}

// ListUsers returns all users.
func (s *Service) ListUsers(ctx context.Context) ([]*domain.User, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("service.ListUsers: %w", err)
	}
	return users, nil
}

// CreateUser validates input and persists a new user.
func (s *Service) CreateUser(ctx context.Context, cmd domain.CreateCommand) (*domain.User, error) {
	if !domain.IsValidRole(cmd.Role) {
		return nil, domain.ErrInvalidRole
	}

	u := &domain.User{
		FirstName: cmd.FirstName,
		LastName:  cmd.LastName,
		Email:     cmd.Email,
		Role:      cmd.Role,
	}

	created, err := s.repo.Save(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("service.CreateUser: %w", err)
	}
	return created, nil
}

// UpdateUser applies a partial update to an existing user.
func (s *Service) UpdateUser(ctx context.Context, id int64, cmd domain.UpdateCommand) (*domain.User, error) {
	if cmd.Role != nil && !domain.IsValidRole(*cmd.Role) {
		return nil, domain.ErrInvalidRole
	}

	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service.UpdateUser: %w", err)
	}

	if cmd.FirstName != nil {
		existing.FirstName = *cmd.FirstName
	}
	if cmd.LastName != nil {
		existing.LastName = *cmd.LastName
	}
	if cmd.Email != nil {
		existing.Email = *cmd.Email
	}
	if cmd.Role != nil {
		existing.Role = *cmd.Role
	}

	updated, err := s.repo.Update(ctx, existing)
	if err != nil {
		return nil, fmt.Errorf("service.UpdateUser: %w", err)
	}
	return updated, nil
}

// DeleteUser removes a user by ID.
func (s *Service) DeleteUser(ctx context.Context, id int64) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("service.DeleteUser: %w", err)
	}
	return nil
}
