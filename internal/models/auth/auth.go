package auth

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r AuthRequest) Validate() error {
	if strings.TrimSpace(r.Email) == "" {
		return errors.New("email is required")
	}

	if strings.TrimSpace(r.Password) == "" {
		return errors.New("password is required")
	}

	return nil
}

type AuthResponse struct {
	Token string `json:"token"`
}

type AuthenticatedUser struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
}
