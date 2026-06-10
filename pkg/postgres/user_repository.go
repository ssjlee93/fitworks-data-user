package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	domain "github.com/ssjlee93/fitworks-data-user/core/domain/user"
)

const pgUniqueViolation = "23505"

// UserRepository implements domain.Repository against PostgreSQL.
type UserRepository struct {
	pool *pgxpool.Pool
}

// New constructs a UserRepository.
func New(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

// FindByID retrieves a user by primary key.
func (r *UserRepository) FindByID(ctx context.Context, id int64) (*domain.User, error) {
	const q = `
		SELECT id, first_name, last_name, email, role, created_at, updated_at
		FROM users
		WHERE id = $1`

	u, err := scanUser(r.pool.QueryRow(ctx, q, id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("postgres.FindByID: %w", err)
	}
	return u, nil
}

// FindAll retrieves every user ordered by id.
func (r *UserRepository) FindAll(ctx context.Context) ([]*domain.User, error) {
	const q = `
		SELECT id, first_name, last_name, email, role, created_at, updated_at
		FROM users
		ORDER BY id`

	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("postgres.FindAll: %w", err)
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		u, err := scanUser(rows)
		if err != nil {
			return nil, fmt.Errorf("postgres.FindAll scan: %w", err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("postgres.FindAll rows: %w", err)
	}
	return users, nil
}

// Save inserts a new user and returns the persisted record (with generated id/timestamps).
func (r *UserRepository) Save(ctx context.Context, u *domain.User) (*domain.User, error) {
	const q = `
		INSERT INTO users (first_name, last_name, email, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id, first_name, last_name, email, role, created_at, updated_at`

	created, err := scanUser(r.pool.QueryRow(ctx, q, u.FirstName, u.LastName, u.Email, string(u.Role)))
	if err != nil {
		if isUniqueViolation(err) {
			return nil, domain.ErrEmailConflict
		}
		return nil, fmt.Errorf("postgres.Save: %w", err)
	}
	return created, nil
}

// Update persists changes to an existing user and returns the updated record.
func (r *UserRepository) Update(ctx context.Context, u *domain.User) (*domain.User, error) {
	const q = `
		UPDATE users
		SET first_name = $1,
		    last_name  = $2,
		    email      = $3,
		    role       = $4,
		    updated_at = NOW()
		WHERE id = $5
		RETURNING id, first_name, last_name, email, role, created_at, updated_at`

	updated, err := scanUser(r.pool.QueryRow(ctx, q, u.FirstName, u.LastName, u.Email, string(u.Role), u.ID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		if isUniqueViolation(err) {
			return nil, domain.ErrEmailConflict
		}
		return nil, fmt.Errorf("postgres.Update: %w", err)
	}
	return updated, nil
}

// Delete removes a user by id.
func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	const q = `DELETE FROM users WHERE id = $1`

	tag, err := r.pool.Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("postgres.Delete: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}

// ── helpers ──────────────────────────────────────────────────────────────────

// scanner abstracts pgx.Row and pgx.Rows so scanUser works for both.
type scanner interface {
	Scan(dest ...any) error
}

func scanUser(s scanner) (*domain.User, error) {
	var u domain.User
	var role string
	err := s.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&role,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	u.Role = domain.Role(role)
	return &u, nil
}

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == pgUniqueViolation
}
