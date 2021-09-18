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
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

}
