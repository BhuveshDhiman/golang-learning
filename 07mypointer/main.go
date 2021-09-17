package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	// var ptr *int
	// fmt.Println("THe value of ptr is ", ptr)

	myNumber := 32
	var ptr2 = &myNumber

	fmt.Println("The address of myNumber is ", ptr2)
	fmt.Println("The value of myNumber is ", *ptr2)

	*ptr2 = *ptr2 * 2
	fmt.Println("New myNumber value is ", myNumber)
}
