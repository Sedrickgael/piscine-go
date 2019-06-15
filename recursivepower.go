package piscine


func RecursivePower(nbr int, power int) int {
	if power == 0{
		return 1
	}
	if power>0 {
		return nbr * RecursivePower(nbr , (power-1))
	}
	if (power < 0){
		return 0
	}

	return 0
}
