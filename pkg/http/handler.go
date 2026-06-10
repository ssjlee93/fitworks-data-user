package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	domain "github.com/ssjlee93/fitworks-data-user/core/domain/user"
)

// Handler holds the HTTP driving adapter for the user service.
type Handler struct {
	svc domain.Service
}

// NewHandler constructs a Handler.
func NewHandler(svc domain.Service) *Handler {
	return &Handler{svc: svc}
}

// ListUsers handles GET /api/v1/users
func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.svc.ListUsers(r.Context())
	if err != nil {
		respondError(w, err)
		return
	}
	respond(w, http.StatusOK, toResponseList(users))
}

// GetUser handles GET /api/v1/users/{id}
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := pathID(r)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	u, err := h.svc.GetUser(r.Context(), id)
	if err != nil {
		respondError(w, err)
		return
	}
	respond(w, http.StatusOK, toResponse(u))
}

// CreateUser handles POST /api/v1/users
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := decodeJSON(r, &req); err != nil {
		http.Error(w, "malformed JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.svc.CreateUser(r.Context(), domain.CreateCommand{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Role:      req.Role,
	})
	if err != nil {
		respondError(w, err)
		return
	}
	respond(w, http.StatusCreated, toResponse(u))
}

// UpdateUser handles PUT /api/v1/users/{id} (full) and PATCH (partial).
// Both verbs use the same handler since UpdateCommand already supports partial fields.
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := pathID(r)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var req UpdateUserRequest
	if err := decodeJSON(r, &req); err != nil {
		http.Error(w, "malformed JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.svc.UpdateUser(r.Context(), id, domain.UpdateCommand{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Role:      req.Role,
	})
	if err != nil {
		respondError(w, err)
		return
	}
	respond(w, http.StatusOK, toResponse(u))
}

// DeleteUser handles DELETE /api/v1/users/{id}
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := pathID(r)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	if err := h.svc.DeleteUser(r.Context(), id); err != nil {
		respondError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// ── helpers ──────────────────────────────────────────────────────────────────

// pathID extracts and parses the {id} path value from the request.
func pathID(r *http.Request) (int64, error) {
	raw := r.PathValue("id")
	id, err := strconv.ParseInt(raw, 10, 64)
	if err != nil || id < 1 {
		return 0, fmt.Errorf("invalid id: %q", raw)
	}
	return id, nil
}

// decodeJSON reads and decodes the request body into dst, rejecting unknown fields.
func decodeJSON(r *http.Request, dst any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(dst)
}
