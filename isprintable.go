package piscine

import "unicode"

func IsPrintable(str string) bool {
	for _, isp := range str {
		if !unicode.IsPrint(isp) {
			return false
		}
	}
	return true
}
