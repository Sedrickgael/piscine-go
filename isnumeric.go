package piscine

import "regexp"

func IsNumeric(str string) bool {
	isn := regexp.MustCompile("^[0-9_]*$")
	return isn.MatchString(str)
}
