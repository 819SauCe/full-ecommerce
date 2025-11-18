package auth

import "full-ecommerce/internal/config"

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

func GetUserIDAndRoleByEmail(email string) (string, string, error) {
	query := `SELECT id, role FROM users WHERE email = $1`
	row := config.DB.QueryRow(query, email)

	var id string
	var role string

	if err := row.Scan(&id, &role); err != nil {
		return "", "", err
	}

	return id, role, nil
}
