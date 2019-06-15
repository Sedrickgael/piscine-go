package piscine

func RecursiveFactorial(nbr int) int {
	if nbr == 1 || nbr == 0{
		return 1
	}
	if nbr > 1 {
		return nbr * RecursiveFactorial(nbr - 1)
	}
	return 0
}

