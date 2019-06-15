package piscine

func IterativeFactorial(nbr int) int {
	fac := 1
	if nbr<0 {
		fac=0
	}
	if nbr==0 || nbr == 1 {
		fac = 1
	}
	for i:=nbr; i>1; i-- {
		fac *= i
	}	
	return fac
}
