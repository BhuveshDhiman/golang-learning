package main

import "fmt"

func main() {
	fmt.Println("If else in Golang")

	loginCount := 7
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount < 20 {
		result = "Old user"
	} else {
		result = "Legend"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("Not less than 10")
	}

	// if err!= nil{

	// }
}
