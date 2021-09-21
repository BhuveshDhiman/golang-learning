package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// No inheritance in golang
	// No super of parent
	bhuvesh := User{"Bhuvesh", "bhuvesh@google.com", true, 21}
	fmt.Println(bhuvesh)
	fmt.Printf("Bhuvesh details are: %+v\n", bhuvesh)
	fmt.Printf("Bhuvesh age is: %v\n", bhuvesh.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
