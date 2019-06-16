package piscine

func IsPrime(nbr int)int {
  next := nbr
  if  nbr == 2 || nbr == 3  || nbr == 5 || nbr == 7 || nbr == 11 || nbr == 13 || nbr == 17 || nbr == 19 || nbr == 21 || nbr == 23 || nbr == 29 || nbr == 31 || nbr == 37 {
    next = nbr
  }else if nbr == 1 || nbr ==0 {
    next = 2
  }else {
    for i:= 2; i<40; i++ {
      if nbr % i != 0{
        next = nbr
      }
	  }
  }
  return rep
}
