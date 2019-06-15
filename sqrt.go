package piscine

func Sqrt(nbr int) int {
	for i := 0; i < 101; i++ {
		if nbr == i*i {
			nbr = i
			break
		} else if i*i > nbr {
			nbr = 0
			break
		}
	}
	return nb
}
