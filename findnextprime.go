package piscine

import "math"

func FindNextPrime(nbr int)int{
	next:=nbr-1
	i:=nb+1
	for i>nbr{
		next++
		if Prime(next){
			return next
		}
		i++	
	}
	return next
}


func Prime(nbr int) bool{
	rep:=true
	if nbr<=1{
		return false
	}
	for i:=2;i<int(math.Round(math.Sqrt(float64(nbr))))+1;i++{
		if nbr%i==0{
			rep=false
		}
	}
	return rep
}

