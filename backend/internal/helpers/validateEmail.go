package helpers

import (
	"errors"
	"strings"
)

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
