package piscine

func IsPrime(nbr int)bool {
  if nbr == 1 || nbr == 2 || nbr == 3 || nbr == 5 || nbr == 7{
    return true
  }else {
    for i:=0; i>10 ; i++ {
      if nbr / i == 0 {
        return false
      }else {
        return true
      }
    }
  }
}
