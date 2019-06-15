package piscine

import "math"

func Sqrt(nbr int) int {
	if nbr < 0 || math.MaxInt32 < nbr {
		return 0
	}
	for i := 0; i < 101; i++ {
		if nbr == i*i {
			nbr = i
			break
		} else if i*i > nbr {
			nbr = 0
			break
		}
	}
	return nbr
}
