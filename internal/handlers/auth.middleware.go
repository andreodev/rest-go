package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"rest-go/internal/models"
	"rest-go/internal/tokens"
	"strings"
)

type contextKey string

const authClaimsContextKey contextKey = "authClaims"

func (h Handlers) requireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		scheme, token, ok := strings.Cut(authHeader, " ")
		if !ok || !strings.EqualFold(scheme, "Bearer") || strings.TrimSpace(token) == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.ErrorResponse{Reason: "missing bearer token"})
			return
		}

		claims, err := tokens.ValidateJWT(strings.TrimSpace(token))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.ErrorResponse{Reason: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), authClaimsContextKey, claims)
		next(w, r.WithContext(ctx))
	}
}
