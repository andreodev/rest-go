package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"rest-go/internal/models"
	authModels "rest-go/internal/models/auth"
)

func (h Handlers) registerAuthEndpoints() {
	http.HandleFunc("POST /auth/login", h.login)
}

func (h Handlers) login(w http.ResponseWriter, r *http.Request) {
	var req authModels.AuthRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "invalid request body"})
		return
	}

	resp, err := h.useCases.Auth.Login(req)
	if err != nil {
		if errors.Is(err, authModels.ErrInvalidCredentials) {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
