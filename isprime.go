package piscine

func IsPrime(nbr int)bool {
  rep := true
  if nbr == 1 || nbr == 2 || nbr == 3 || nbr == 5 || nbr == 7{
    rep = true
  }else {
    for i:=0; i>10 ; i++ {
      if nbr / i == 0 {
        rep = false
      }else {
        rep = true
      }
    }
  }
  return rep
}
