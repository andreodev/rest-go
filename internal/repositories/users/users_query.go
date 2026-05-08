package users

import (
	"database/sql"
	userModels "rest-go/internal/models/users"
)

type Users struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Users {
	return &Users{db: db}
}

func (u Users) GetAll() ([]userModels.User, error) {
	rows, err := u.db.Query(`
		SELECT id, name, email, password
		FROM users
		ORDER BY created_at ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]userModels.User, 0)

	for rows.Next() {
		var user userModels.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u Users) EmailExist(email string) (bool, error) {
	var exists bool
	err := u.db.QueryRow(`
		SELECT EXISTS(
			SELECT 1
			FROM users
			WHERE email = $1
		)
	`, email).Scan(&exists)

	return exists, err
}

func (u *Users) Add(newUser userModels.User) error {
	_, err := u.db.Exec(`
		INSERT INTO users (id, name, email, password)
		VALUES ($1, $2, $3, $4)
	`, newUser.ID, newUser.Name, newUser.Email, newUser.Password)

	return err
}

func (u *Users) DeleteById(id string) error {
	result, err := u.db.Exec(`
		DELETE FROM users
		WHERE id = $1
	`, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (u Users) GetById(id string) (userModels.User, error) {
	var user userModels.User
	err := u.db.QueryRow(`
		SELECT id, name, email, password
		FROM users
		WHERE id = $1
	`, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	return user, err
}
