package piscine

func IsPrime(nbr int)bool {
  if  nbr == 2 || nbr == 3  || nbr == 5 || nbr == 7  {
    return true
  }else if nbr == 1 {
    return false
  }else {
    for i:= 2; i<100; i++ {
      if nbr % i == 0{
        return false
      }else {
        return true
      }
    }
  }
}
