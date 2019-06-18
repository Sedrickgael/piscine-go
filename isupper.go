package piscine

import "unicode"

func IsUpper(str string) bool {
	for _, isu := range str {
		if !unicode.IsUpper(isu) {
			return false
		}
	}
	return true
}
