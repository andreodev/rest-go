package users

import (
	"database/sql"
	"rest-go/internal/models"
)

type Users struct {
	users []models.User
	db    *sql.DB
}

func NewMemory() *Users {
	return &Users{users: make([]models.User, 0)}
}

func NewPostgres(db *sql.DB) *Users {
	return &Users{db: db}
}

func (u Users) GetAll() ([]models.User, error) {
	if u.db == nil {
		return u.users, nil
	}

	rows, err := u.db.Query(`
		SELECT id, name, email
		FROM users
		ORDER BY created_at ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]models.User, 0)

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
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
	if u.db != nil {
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

	for _, user := range u.users {
		if user.Email == email {
			return true, nil
		}
	}

	return false, nil
}

func (u *Users) Add(newUser models.User) error {
	if u.db != nil {
		_, err := u.db.Exec(`
			INSERT INTO users (id, name, email)
			VALUES ($1, $2, $3)
		`, newUser.ID, newUser.Name, newUser.Email)

		return err
	}

	u.users = append(u.users, newUser)
	return nil
}
