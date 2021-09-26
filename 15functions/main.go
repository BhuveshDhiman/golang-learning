package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)

	fmt.Println("Result is: ", result)

	proResult, msg := proAdder(2, 5, 1, 2)
	fmt.Println("Pro result is: ", proResult, msg)
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hi pro result!"
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func greeter() {
	fmt.Println("Hello from Golang")
}
