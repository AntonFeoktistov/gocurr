package service

import "unicode/utf8"

func validateCode(code string) bool {
	countRunes := utf8.RuneCountInString(code)
	return countRunes == 3
}
