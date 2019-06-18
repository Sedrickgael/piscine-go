package piscine

import "strings"

func Capitalize(word string) string {
	cap := strings.Title(ToLower(word))
	i := Index(word, "_")
	if i == -1 {
		return cap
	} else {
		cap = strings.Replace(cap, string(cap[i+1]), ToUpper(string(cao[i+1])), i+1)
		return cap
	}
}
