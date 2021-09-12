package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	// fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006 03:04 pm Monday"))

	createdDate := time.Date(2000, time.July, 10, 11, 15, 0, 0, time.Local)
	fmt.Println(createdDate.Format("02-01-2006 03:04 pm Monday"))
}
