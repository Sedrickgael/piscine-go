package piscine

import "fmt"

func IterativeFactoriel( nbr int) int {
		fac := 1
		for i := nbr; i > 1; i--{
			fac *= i
		}	
		fmt.Print(fac)
}
