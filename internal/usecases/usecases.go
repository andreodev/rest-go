package usecases

import (
	"errors"
	"log/slog"
	"rest-go/internal/models"
	"rest-go/internal/repositories"

	"github.com/google/uuid"
)

type UseCases struct {
	repos *repositories.Repositories
}

func New(repos *repositories.Repositories) *UseCases {
	return &UseCases{repos: repos}
}

func (u UseCases) GetAllUsers() []models.User {
	users := u.repos.User.GetAll()

	return users
}

func (u UseCases) AddUser(newUser models.UserCreateRequest) (uuid.UUID, error) {

	exist := u.repos.User.EmailExist(newUser.Email)

	if exist {
		slog.Error("this users already exist", "email", newUser.Email)
		return uuid.Nil, errors.New("email already exist")
	}

	repoReq := models.User{
		ID:    uuid.New(),
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	u.repos.User.Add(repoReq)

	return repoReq.ID, nil
}
