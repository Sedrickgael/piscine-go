package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {

	args := os.Args

	sort.Strings(args)

	for i := 0; i < len(args)-1; i++ {
		fmt.Println(args[i])
	}
}
