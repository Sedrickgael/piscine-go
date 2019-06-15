package piscine

import "math"

func RecursiveFactorial(nbr int) int {
	if nbr < 0 {
		return 0
	} else if nbr == 0 || nbr == 1 {
		return 1
	} else {
		fac := 1
		for i := nbr; i >=1; i-- {
			fac *= i
			if math.MaxInt32 < fac {
				fac = 0
				break
			}
		}
		return fac
	}
}
