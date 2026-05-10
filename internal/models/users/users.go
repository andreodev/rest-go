package users

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password string
}

type UserReponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r UserCreateRequest) Validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errors.New("name is required")
	}

	if strings.TrimSpace(r.Email) == "" {
		return errors.New("email is required")
	}

	if strings.TrimSpace(r.Password) == "" {
		return errors.New("password is required")
	}

	return nil
}

type UserCreateResponse struct {
	NewUserID uuid.UUID `json:"newUserId"`
}

type UserDeleteRequest struct {
	Id uuid.UUID `json:"id"`
}

type UserDeleteResponse struct {
	Message string    `json:"message"`
	ID      uuid.UUID `json:"id"`
}
