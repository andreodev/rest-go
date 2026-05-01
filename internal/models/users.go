package models

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID
	Name  string
	Email string
}

type UserCreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
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
