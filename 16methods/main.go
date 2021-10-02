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

	bhuvesh.GetStatus()
	bhuvesh.NewMail()
	fmt.Printf("Bhuvesh details are: %+v\n", bhuvesh)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "newEmail@google.com"
	fmt.Println("Email of this user is: ", u.Email)
}
