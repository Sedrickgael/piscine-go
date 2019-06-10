package piscine

import "fmt"
import "strconv"

func PrintComb2()  {
	for i:=0;i<10;i++{
		for j:=0;j<10;j++{
			for k:=0;k<10;k++{
				for l:=0;l<10;l++{
					nbr1:= strconv.Itoa(i) + strconv.Itoa(j)
					nbr2:= strconv.Itoa(k) + strconv.Itoa(l)
					if(nbr1 < nbr2){
						fmt.Print(nbr1 + " " +nbr2)
						
						if nbr1 < "98" {
							fmt.Print(", ")
						}else {
							fmt.Print("\n")
						}
					}
				}
			}
		}
	}
}
