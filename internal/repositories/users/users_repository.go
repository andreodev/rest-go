package users

import userModels "rest-go/internal/models/users"

type UserRepository interface {
	DeleteById(id string) error
	GetAll() ([]userModels.User, error)
	GetById(id string) (userModels.User, error)
	Add(newUser userModels.User) error
	EmailExist(email string) (bool, error)
}
