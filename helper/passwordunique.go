package helper

import "unicode"

func IsPasswordValid(password string) bool {
	var (
		hasUpper   bool
		hasLower   bool
		hasDigit   bool
		hasSpecial bool
	)

	for _, r := range password {
		switch {
		case unicode.IsUpper(r): 
				hasUpper = true
		case unicode.IsLower(r): 
				hasLower = true
		case unicode.IsDigit(r): 
				hasDigit = true 
		case unicode.IsSymbol(r) || unicode.IsPunct(r):
				hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}