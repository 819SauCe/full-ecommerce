package helpers

import (
	"errors"
	"strings"
	"unicode"
)

func PasswordIsValid(password string) (bool, error) {
	if len(password) < 6 {
		return false, errors.New("password must be at least 6 characters long")
	}

	if len(password) > 128 {
		return false, errors.New("password must be at most 128 characters long")
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasDigit   bool
		hasSpecial bool
	)

	//Check if all characters are the same (e.g., "aaaaaaaaaaaa")
	allSame := true
	var firstRune rune
	for i, r := range password {
		if i == 0 {
			firstRune = r
		} else if r != firstRune {
			allSame = false
		}

		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		case unicode.IsSpace(r):
			return false, errors.New("password cannot contain spaces")
		}
	}

	if allSame {
		return false, errors.New("password cannot be the same character repeated")
	}

	if !hasUpper {
		return false, errors.New("password must contain at least one uppercase letter")
	}

	if !hasLower {
		return false, errors.New("password must contain at least one lowercase letter")
	}

	if !hasDigit {
		return false, errors.New("password must contain at least one digit")
	}

	if !hasSpecial {
		return false, errors.New("password must contain at least one special character")
	}

	return true, nil
}

func NameIsValid(name string) (bool, error) {
	number_regex := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	specialCharacters := []string{"!", "@", "#", "$", "%", "¨", "&", "*", "(", ")", "-", "+", "=", ",", ".", "'", "`", "´", "/", "\\", "|"}

	if name == "" {
		return false, errors.New("the name cannot be null")
	}

	for i := 0; i < len(number_regex); i++ {
		if strings.Contains(name, number_regex[i]) {
			return false, errors.New("name can not have numbers")
		}
	}

	for i := 0; i < len(number_regex); i++ {
		if strings.Contains(name, specialCharacters[i]) {
			return false, errors.New("name can not have special caracterers")
		}
	}

	if len(name) > 100 {
		return false, errors.New("the name cannot be longer than 200 characters")
	}

	if len(name) < 3 {
		return false, errors.New("the name cannot be short than 3 characters")
	}

	return true, nil
}

func EmailIsValid(email string) (bool, error) {
	if email == "" {
		return false, errors.New("the email address cannot be empty")
	}
	if len(email) > 200 {
		return false, errors.New("the email cannot be longer than 200 characters")
	}
	if len(email) < 12 {
		return false, errors.New("the email cannot be short than 3 characters")
	}
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return false, errors.New("email is not valid")
	}
	if strings.Contains(email, `"`) || strings.Contains(email, "'") {
		return false, errors.New("the email cannot contain quotation marks")
	}
	if strings.Contains(email, "(") || strings.Contains(email, ")") {
		return false, errors.New("the email address cannot contain square brackets")
	}
	if strings.Contains(email, "{") || strings.Contains(email, "}") {
		return false, errors.New("the email cannot contain keys")
	}
	if strings.Contains(email, "!") || strings.Contains(email, "#") || strings.Contains(email, "$") || strings.Contains(email, "%") || strings.Contains(email, "&") || strings.Contains(email, "*") || strings.Contains(email, "¨") {
		return false, errors.New("the email cannot contain special characters")
	}

	return true, nil
}
