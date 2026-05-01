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

func (u UseCases) GetAllUsers() ([]models.User, error) {
	users, err := u.repos.User.GetAll()
	if err != nil {
		slog.Error("failed to get users", "err", err)
		return nil, err
	}

	return users, nil
}

func (u UseCases) AddUser(newUser models.UserCreateRequest) (uuid.UUID, error) {

	exist, err := u.repos.User.EmailExist(newUser.Email)
	if err != nil {
		slog.Error("failed to check user email", "email", newUser.Email, "err", err)
		return uuid.Nil, err
	}

	if exist {
		slog.Error("this users already exist", "email", newUser.Email)
		return uuid.Nil, errors.New("email already exist")
	}

	repoReq := models.User{
		ID:    uuid.New(),
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	if err := u.repos.User.Add(repoReq); err != nil {
		slog.Error("failed to add user", "email", newUser.Email, "err", err)
		return uuid.Nil, err
	}

	return repoReq.ID, nil
}

func (u UseCases) DeleteUserById(id string) error {
	if err := u.repos.User.DeleteById(id); err != nil {
		slog.Error("failed to delete user", "id", id, "err", err)
		return err
	}

	return nil
}

func (u UseCases) GetUserById(id string) (models.User, error) {
	user, err := u.repos.User.GetById(id)
	if err != nil {
		slog.Error("failed to get user by id", "id", id, "err", err)
		return models.User{}, err
	}

	return user, nil
}
