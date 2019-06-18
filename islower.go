package piscine

import "unicode"

func IsLower(str string) bool {
	for _, isl := range str {
		if !unicode.IsLower(isl) {
			return false
		}
	}
	return true
}
