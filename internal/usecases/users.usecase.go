package usecases

import (
	"errors"
	"log/slog"
	userModels "rest-go/internal/models/users"
	userRepo "rest-go/internal/repositories/users"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UsersUseCase struct {
	repo userRepo.UserRepository
}

func NewUsersUseCase(repo userRepo.UserRepository) *UsersUseCase {
	return &UsersUseCase{repo: repo}
}

func (u UsersUseCase) GetAll() ([]userModels.User, error) {
	users, err := u.repo.GetAll()
	if err != nil {
		slog.Error("failed to get users", "err", err)
		return nil, err
	}

	return users, nil
}

func (u UsersUseCase) Add(newUser userModels.UserCreateRequest) (uuid.UUID, error) {
	exist, err := u.repo.EmailExist(newUser.Email)
	if err != nil {
		slog.Error("failed to check user email", "email", newUser.Email, "err", err)
		return uuid.Nil, err
	}

	if exist {
		slog.Error("this users already exist", "email", newUser.Email)
		return uuid.Nil, errors.New("email already exist")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(newUser.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		slog.Error("failed to hash password", "email", newUser.Email, "err", err)
		return uuid.Nil, err
	}

	repoReq := userModels.User{
		ID:       uuid.New(),
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: string(hash),
	}

	if err := u.repo.Add(repoReq); err != nil {
		slog.Error("failed to add user", "email", newUser.Email, "err", err)
		return uuid.Nil, err
	}

	return repoReq.ID, nil
}

func (u UsersUseCase) DeleteById(id string) error {
	if err := u.repo.DeleteById(id); err != nil {
		slog.Error("failed to delete user", "id", id, "err", err)
		return err
	}

	return nil
}

func (u UsersUseCase) GetById(id string) (userModels.User, error) {
	user, err := u.repo.GetById(id)
	if err != nil {
		slog.Error("failed to get user by id", "id", id, "err", err)
		return userModels.User{}, err
	}

	return user, nil
}
