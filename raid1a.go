package main

import "fmt"

func Raid1a(x, y int){
	for i:=1; i<=x; i++{
		if i==1 || i==x{
			fmt.Print("o")
		}else {
			fmt.Print("-")
		}		
	}
	if y>1 {
		fmt.Print("\n")
	}
	
	for j:=2; j<=y-1 ; j++{
		if x == 1 {
			fmt.Println("|")
		}else {
			fmt.Print("|")
			for i:=1; i<=x-2; i++{
				fmt.Print(" ")	
			}
			fmt.Println("|")
		}
		
		
	}
	if y>1{
		for i:=1; i<=x; i++{
			if i==1 || i==x{
				fmt.Print("0")
			}else {
				fmt.Print("-")
			}		
		}
	}
	
}