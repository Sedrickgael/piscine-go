package main

import (
		"fmt"
    "os"
)

func main(){
	args := os.Args[1:]
	for _,params:= range args{
		fmt.Println(params)
	}
}
