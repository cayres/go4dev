package main

import (
	"unicode"
)

func isValidPassword(password string) bool {
	passwordLength := len(password)
	atLeastOneUppercaseLetter := false
	atLeastOneDigit := false

	if passwordLength < 5 || passwordLength > 12 {
		return false
	}

	for _, char := range password {

		if unicode.IsUpper(char) && unicode.IsLetter(char) {
			atLeastOneUppercaseLetter = true
		}

		if unicode.IsDigit(char) {
			atLeastOneDigit = true
		}
	}

	return atLeastOneDigit && atLeastOneUppercaseLetter
}
