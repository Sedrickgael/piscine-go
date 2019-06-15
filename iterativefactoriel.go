package piscine

func IterativeFactorial(nbr int) int {
	
	if nbr<0 {
		return 0
	}
	if nbr==0 || nbr == 1 {
		return 1
	}else {
		fac := 1
		for i:=nbr; i>1; i-- {
			fac *= i
		}
		return fac
	}
	
}
