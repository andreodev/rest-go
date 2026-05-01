package repositories

import (
	"database/sql"
	"rest-go/internal/models"
	"rest-go/internal/repositories/users"
)

type Repositories struct {
	User interface {
		DeleteById(id string) error
		GetAll() ([]models.User, error)
		GetById(id string) (models.User, error)
		Add(newUser models.User) error
		EmailExist(email string) (bool, error)
	}
}

func New(db *sql.DB) *Repositories {
	return &Repositories{
		User: users.NewPostgres(db),
	}
}
