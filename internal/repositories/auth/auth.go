package auth

import (
	"database/sql"
	authModels "rest-go/internal/models/auth"
)

type Auth struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Auth {
	return &Auth{db: db}
}

func (a *Auth) FindByEmail(email string) (authModels.AuthenticatedUser, error) {
	var user authModels.AuthenticatedUser

	err := a.db.QueryRow(`
		SELECT id, email, password 
		FROM users
		WHERE email = $1 
	`, email).Scan(&user.ID, &user.Email, &user.PasswordHash)

	return user, err
}
