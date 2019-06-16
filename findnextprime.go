package piscine

func Prime(nbr int)bool {
  rep := true
  if  nbr == 2 || nbr == 3  || nbr == 5 || nbr == 7 || nbr == 11 || nbr == 13 || nbr == 17 || nbr == 19 || nbr == 21 || nbr == 23 || nbr == 29 || nbr == 31 || nbr == 37 {
    rep = true
  }else if nbr == 1 || nbr ==0 {
    rep = false
  }else {
    for i:= 2; i<40; i++ {
      if nbr % i == 0{
        rep = false
      }
	}
	
  }
  return rep
}


func FindNextPrime(nbr int)int {
	next := nbr 
	for i:= nbr; i>nbr;  {
		if Prime(i){
			next=i
		}else {
			i++
		}
	}
	return next
}
