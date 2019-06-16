package piscine

func IsPrime(nbr int)bool {
  rep := true
  if  nbr % 2 == 0|| nbr % 3 ==0 || nbr % 5 ==0 || nbr % 7== 0{
    rep = false
  }else {
   rep = true
  }
  return rep
}
