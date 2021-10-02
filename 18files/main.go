package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")

	content := "This needs to go in a file"

	file, err := os.Create("./myFile.txt")
	checkNilErr(err)
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length is ", length)
	defer file.Close()
	readFile("./myFile.txt")

}

func readFile(fileName string) {
	byteData, err := ioutil.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println("Text data inside the file is \n", string(byteData))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
