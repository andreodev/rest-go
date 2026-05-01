package repositories

import (
	"database/sql"
	"rest-go/internal/models"
	"rest-go/internal/repositories/users"
)

type Repositories struct {
	User interface {
		GetAll() ([]models.User, error)
		Add(newUser models.User) error
		EmailExist(email string) (bool, error)
	}
}

func New(db *sql.DB) *Repositories {
	if db != nil {
		return &Repositories{
			User: users.NewPostgres(db),
		}
	}

	return &Repositories{
		User: users.NewMemory(),
	}
}
