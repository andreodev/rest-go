package repositories

import (
	"database/sql"
	"rest-go/internal/repositories/products"
	"rest-go/internal/repositories/users"
)

type Repositories struct {
	User     users.UserRepository
	Products products.ProductRepository
}

func New(db *sql.DB) *Repositories {
	return &Repositories{
		User:     users.NewPostgres(db),
		Products: products.NewPostgres(db),
	}
}
