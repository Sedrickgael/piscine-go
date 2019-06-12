package piscine

import "fmt"

func PointOne (a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
	fmt.Print(div)
	fmt.Print(mod)
}
