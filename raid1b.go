package piscine

import "fmt"

func Raid1b(x, y int){
  if x>0 && y>0 {
    for i:=1; i<=x; i++{
      if i==1 {
        fmt.Print("/")
      }else if i==x{
        fmt.Print("\\")
      }else {
        fmt.Print('*')
      }

    }
    fmt.Print("\n")

    for j:=2; j<=y-1 ; j++{
      if x == 1 {
        fmt.Println("*")
      }else {
        fmt.Print("*")
        for i:=1; i<=x-2; i++{
          fmt.Print(" ")	
        }
        fmt.Println("*")
      }
    }
    if y>1{
      for i:=1; i<=x; i++{
        if i==1 {
          fmt.Print("/")
        }else if i==x{
          fmt.Print("\\")
        }else {
          fmt.Print('*')
        }

      }
    }
   }	
}
