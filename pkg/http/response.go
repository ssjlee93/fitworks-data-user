package http

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	domain "github.com/ssjlee93/fitworks-data-user/core/domain/user"
)

// UserResponse is the JSON representation of a User returned to callers.
type UserResponse struct {
	ID        int64       `json:"id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `json:"email"`
	Role      domain.Role `json:"role"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
}

func toResponse(u *domain.User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: u.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func toResponseList(users []*domain.User) []UserResponse {
	out := make([]UserResponse, len(users))
	for i, u := range users {
		out[i] = toResponse(u)
	}
	return out
}

// envelope is the standard JSON response wrapper.
type envelope struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func respond(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(envelope{Data: data}); err != nil {
		slog.Error("respond: encode", "err", err)
	}
}

func respondError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	msg := "internal server error"

	switch {
	case errors.Is(err, domain.ErrNotFound):
		status = http.StatusNotFound
		msg = "user not found"
	case errors.Is(err, domain.ErrEmailConflict):
		status = http.StatusConflict
		msg = "email already in use"
	case errors.Is(err, domain.ErrInvalidRole):
		status = http.StatusUnprocessableEntity
		msg = err.Error()
	}

	slog.Error("request error", "status", status, "err", err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(envelope{Error: msg})
}
