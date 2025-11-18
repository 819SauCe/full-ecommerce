package auth

import (
	"fmt"
	"full-ecommerce/internal/helpers"
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

	//validate first_name
	isValid, err = helpers.NameIsValid(input.First_name)
	if err != nil {
		return err
	}
	if !isValid {
		return fmt.Errorf("first_name is not valid")
	}

	//validate last_name
	isValid, err = helpers.NameIsValid(input.Last_name)
	if err != nil {
		return err
	}
	if !isValid {
		return fmt.Errorf("last_name is not valid")
	}

	//Validate email
	isValid, err = helpers.EmailIsValid(input.Email)
	if err != nil {
		return err
	}
	if !isValid {
		return fmt.Errorf("email is not valid")
	}
	if UserAlrdExists(input.Email) {
		return fmt.Errorf("email already exists")
	}

	//Validate password
	isValid, err = helpers.PasswordIsValid(input.Password)
	if err != nil {
		return err
	}
	if !isValid {
		return fmt.Errorf("password is not valid")
	}

	//hash password
	password_hash, err := helpers.HashPassword(input.Password)
	if err != nil {
		return fmt.Errorf("error while hashing password")
	}

	//insert user
	if err := RegisterUser(input.First_name, input.Last_name, input.Email, password_hash); err != nil {
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
	helpers.CheckPasswordHash(input.Password, password_hash)

	return nil
}
