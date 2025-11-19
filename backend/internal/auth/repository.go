package auth

import (
	"database/sql"
	"errors"
	"full-ecommerce/internal/config"
	"full-ecommerce/internal/helpers"
)

func UserAlrdExists(email string) bool {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	row := config.DB.QueryRow(query, email)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false
	}
	return exists
}

func RegisterUser(first, last, email, password string) error {
	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return err
	}

	role := "user"
	if count == 0 {
		role = "admin"
	}

	query := `INSERT INTO users (first_name, last_name, email, password, role) VALUES ($1, $2, $3, $4, $5)`
	_, err = config.DB.Exec(query, first, last, email, password, role)
	return err
}

func GetUserDataByEmail(email string) (helpers.UserData, error) {
	const query = `
		SELECT 
			id,
			first_name,
			last_name,
			email,
			role
		FROM users
		WHERE email = $1
	`

	row := config.DB.QueryRow(query, email)

	var user helpers.UserData

	err := row.Scan(
		&user.Id,
		&user.First_name,
		&user.Last_name,
		&user.Email,
		&user.Role,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return helpers.UserData{}, ErrUserNotFound
		}
		return helpers.UserData{}, err
	}

	user.Profile_img = ""

	return user, nil
}

func GetPasswordHashByEmail(email string) (string, error) {
	query := `SELECT password FROM users WHERE email = $1`
	row := config.DB.QueryRow(query, email)

	var hash string
	if err := row.Scan(&hash); err != nil {
		return "", err
	}

	return hash, nil
}
