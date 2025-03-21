package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/ssjlee93/fitworks-data-user/dtos"
	"net/http"
)

func marshalResponse[T []dtos.User | *dtos.User](res T, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return fmt.Errorf("could not marshal res: %w", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return nil
}

func unmarshalRequest(w http.ResponseWriter, r *http.Request) (*dtos.User, error) {
	var user *dtos.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return nil, fmt.Errorf("could not unmarshal request: %w", err)
	}
	return user, nil
}
