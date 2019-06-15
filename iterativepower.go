package piscine

import "fmt"

func IterativePower(nbr int, power int)int {
	pow :=1
	if ( (nbr>0 && power>0) || (nbr<0 && power%2 == 0) ){
		for i:=power; i>=1; i--{
			pow *=nbr
		}
		
	} else if (power==0){
		pow = 1
	} else {
		pow = 0
	}
	fmt.Print(pow)
}
