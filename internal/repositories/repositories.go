package repositories

import (
	"rest-go/internal/models"
	"rest-go/internal/repositories/users"
)

type Repositories struct {
	User interface {
		GetAll() []models.User
		Add(newUser models.User)
		EmailExist(email string) bool
	}
}

func New() *Repositories {
	return &Repositories{
		User: users.New(),
	}
}
