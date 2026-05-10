package auth

import (
	"database/sql"
	"errors"
	authModels "rest-go/internal/models/auth"
	authRepo "rest-go/internal/repositories/auth"
	"rest-go/internal/tokens"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	repo authRepo.AuthRepositories
}

func NewAuthUseCase(repo authRepo.AuthRepositories) *AuthUseCase {
	return &AuthUseCase{repo: repo}
}

func (a *AuthUseCase) Login(data authModels.AuthRequest) (authModels.AuthResponse, error) {
	if err := data.Validate(); err != nil {
		return authModels.AuthResponse{}, err
	}

	user, err := a.repo.FindByEmail(data.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return authModels.AuthResponse{}, authModels.ErrInvalidCredentials
		}

		return authModels.AuthResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(data.Password),
	); err != nil {
		return authModels.AuthResponse{}, authModels.ErrInvalidCredentials
	}

	token, err := tokens.GenerateJWT(user.ID.String(), user.Email)
	if err != nil {
		return authModels.AuthResponse{}, err
	}

	return authModels.AuthResponse{Token: token}, nil
}
