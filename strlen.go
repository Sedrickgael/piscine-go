package piscine

import "fmt"

func StrLen (str string) {
	nb := len([]rune(str))
	fmt.Print(nb)
}
