package service

import "unicode/utf8"

func validateCode(code string) bool {
	countRunes := utf8.RuneCountInString(code)
	return countRunes == 3
}

func validateName(name string) bool {
	countRunes := utf8.RuneCountInString(name)
	return countRunes > 0 && countRunes < 25
}

func validateSign(sign string) bool {
	countRunes := utf8.RuneCountInString(sign)
	return countRunes > 0 && countRunes < 3
}
