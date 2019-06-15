package piscine

func IterativeFactorial(nbr int) int {
	fac := 1
	if(nbr<0){
		fac=0
	}
	for i:=nbr; i>=1; i-- {
		fac *= i
	}	
	return fac
}
