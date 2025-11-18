package helpers

import (
	"errors"
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
