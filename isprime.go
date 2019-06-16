package piscine

func IsPrime(nbr int)bool {
  rep := true
  if  nbr == 2 || nbr == 3  || nbr == 5 || nbr == 7  {
    rep = true
  }else if nbr == 1 {
    rep = false
  }else {
    for i:= 2; i<100; i++ {
      if nbr % i == 0{
        rep = false
      }else {
        rep = true
      }
    }
  }
  return rep
}
