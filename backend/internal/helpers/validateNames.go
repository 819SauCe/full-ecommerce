package helpers

import (
	"errors"
	"strings"
)

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
