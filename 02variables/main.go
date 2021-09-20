package main

import "fmt"

const LoginToken string = "qwerty" // capital means public

func main() {
	var username string = "Bhuvesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.54543656323442545545
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// Default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString)

	// implicit type
	var website = "bhuveshdhiman.me"
	fmt.Println(website)

	// no var style or walrus operator
	numberOfUsers := 300000.5
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
