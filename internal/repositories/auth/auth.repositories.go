package auth

import authModel "rest-go/internal/models/auth"

type AuthRepositories interface {
	FindByEmail(email string) (authModel.AuthenticatedUser, error)
}
