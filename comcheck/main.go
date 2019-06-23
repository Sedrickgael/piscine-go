package main

import "fmt"
import "os"

func main(){
	param:=os.Args
	if iscorrectconcheck(param){
		fmt.Println("Alert!!!")	
	}
}

func iscorrectconcheck(param []string) bool{
	if len(param)==1{
		return false
	}else{
		for _,val:= range param[1:]{
			if val=="01" || val=="galaxy" || val=="galaxy 01"{
				return true
			}	
		}
	}
	return false
}
