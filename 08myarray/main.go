package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays!")

	var fruitList [4]string
	fruitList[0] = "Banana"
	fruitList[1] = "Mango"
	fruitList[3] = "Apple"

	fmt.Println("Fruit list is", fruitList)
	fmt.Println("Fruit list length is", len(fruitList))

	var vegList = [3]string{"potato", "beans", "spinach"}
	fmt.Println("Veg list is ", vegList)
	fmt.Println("Veg list length is", len(vegList))
}
