package repositories

import (
	"database/sql"
	userModels "rest-go/internal/models/users"
	"rest-go/internal/repositories/users"
)

type Repositories struct {
	User interface {
		DeleteById(id string) error
		GetAll() ([]userModels.User, error)
		GetById(id string) (userModels.User, error)
		Add(newUser userModels.User) error
		EmailExist(email string) (bool, error)
	}
}

func New(db *sql.DB) *Repositories {
	return &Repositories{
		User: users.NewPostgres(db),
	}
}
