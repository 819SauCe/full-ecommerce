package auth

import (
	"errors"
	"fmt"
	"full-ecommerce/internal/helpers"
)

var (
	ErrInvalidFirstName   = errors.New("invalid_first_name")
	ErrInvalidLastName    = errors.New("invalid_last_name")
	ErrInvalidEmail       = errors.New("invalid_email")
	ErrEmailAlreadyUsed   = errors.New("email_already_used")
	ErrInvalidPassword    = errors.New("invalid_password")
	ErrInvalidCredentials = errors.New("invalid_credentials")
	ErrUserNotFound       = errors.New("user not found")
)

type RegisterModel struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type LoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(input RegisterModel) error {
	var (
		isValid bool
		err     error
	)

	// validate first_name
	isValid, err = helpers.NameIsValid(input.First_name)
	if err != nil {
		return err
	}
	if !isValid {
		return ErrInvalidFirstName
	}

	// validate last_name
	isValid, err = helpers.NameIsValid(input.Last_name)
	if err != nil {
		return err
	}
	if !isValid {
		return ErrInvalidLastName
	}

	// validate email
	isValid, err = helpers.EmailIsValid(input.Email)
	if err != nil {
		return err
	}
	if !isValid {
		return ErrInvalidEmail
	}
	if UserAlrdExists(input.Email) {
		return ErrEmailAlreadyUsed
	}

	// validate password
	isValid, err = helpers.PasswordIsValid(input.Password)
	if err != nil {
		return err
	}
	if !isValid {
		return ErrInvalidPassword
	}

	// hash password
	passwordHash, err := helpers.HashPassword(input.Password)
	if err != nil {
		return fmt.Errorf("error while hashing password: %w", err)
	}

	// insert user
	if err := RegisterUser(input.First_name, input.Last_name, input.Email, passwordHash); err != nil {
		return err
	}

	return nil
}

func Login(input LoginModel) error {
	var (
		isValid bool
		err     error
	)

	//Validate Email
	isValid, err = helpers.EmailIsValid(input.Email)
	if err != nil {
		return err
	}
	if !isValid {
		return fmt.Errorf("email is not valid")
	}
	if !UserAlrdExists(input.Email) {
		return fmt.Errorf("email not exists")
	}

	//Validate Password
	isValid, err = helpers.PasswordIsValid(input.Password)
	if err != nil {
		return err
	}
	if !isValid {
		return fmt.Errorf("password is not valid")
	}

	password_hash, err := GetPasswordHashByEmail(input.Email)
	if err != nil {
		return fmt.Errorf("error")
	}
	passwordIsCorrect := helpers.CheckPasswordHash(input.Password, password_hash)
	if !passwordIsCorrect {
		return ErrInvalidCredentials
	}

	return nil
}

func Logout(sessionToken string) error {
	return nil
}
