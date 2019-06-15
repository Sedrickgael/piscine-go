package piscine

import "fmt"

func RecursiveFactorial(nbr int) int {
	if nbr == 1 || nbr == 0{
		return 1
	}
	if nbr > 1 {
		return nbr * recursivefac(nbr - 1)
	}
	return 0
}

func main(){
	fac := recursivefac(nbr)
	fmt.Print(fac)
}
