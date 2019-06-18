package piscine

import "fmt"

func Raid1e(x, y int){
if x>0 && y>0{
    for i:=1; i<=x; i++{
      if i==1 {
        fmt.Print("A")
      }else if i==x{
        fmt.Print("C")
      }else {
        fmt.Print("B")
      }	
    }
    fmt.Print("\n")
    for j:=2; j<=y-1 ; j++{
      if x == 1 {
        fmt.Println("B")
      }else {
        fmt.Print("B")
        for i:=1; i<=x-2; i++{
          fmt.Print(" ")	
        }
        fmt.Println("B")
      }


    }
    if y>1{
      for i:=1; i<=x; i++{
        if i==1 {
          fmt.Print("C")
        }else if i==x{
          fmt.Print("A")
        }else {
          fmt.Print("B")
        }	
      }
      fmt.Print("\n")
    }
  }
}
