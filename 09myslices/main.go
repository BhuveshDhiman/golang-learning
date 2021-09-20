package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	// var fruitList = []string{"apple", "mango", "watermelon"}
	// fmt.Printf("Type of fruitList is %T\n", fruitList)

	// fruitList = append(fruitList, "Banana", "Tomato")
	// fmt.Println("Fruit list = ", fruitList)

	// fruitList = append(fruitList[:3])
	// fmt.Println("Fruit list = ", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 456
	highScores[3] = 567
	// highScores[4] = 777

	highScores = append(highScores, 555, 666, 333)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	// How to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
