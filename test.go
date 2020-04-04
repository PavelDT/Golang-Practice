package main

import (
	"fmt"
	//"gotest"
)

func main() {
	fmt.Println("this is a GoLang practice")
	var x int

	//reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter a number between 1 and 100")
	fmt.Scan(&x)
	fmt.Print("Your number is", x)

	if x < 50 {
		fmt.Println("LowEnd")
	} else if x > 50 {
		fmt.Println("Highend")
	} else {
		fmt.Println("Invalid Number")
	}

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your city: ")
	// city, _ := reader.ReadString('\n')
	// fmt.Print("You live in " + city)
}
