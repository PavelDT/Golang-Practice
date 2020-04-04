package main

import (
	"fmt"
	//"gotest"
)

func main() {
	fmt.Println("this is a GoLang practice")

	x := 100

	if x == 50 {
		fmt.Println("Germany")
	} else if x == 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}
}
