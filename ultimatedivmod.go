package piscine

import "fmt"

func UltimateDivMod (a *int, b *int) {
	c := *a
	*a = *a / *b
	*b = c % *b
	fmt.Print(a)
	fmt.Print(b)
}
