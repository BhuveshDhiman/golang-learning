package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web requests")

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close() // caller's responsibility to close the connection
	fmt.Printf("Response type = %T\n ", res)

	byteData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	content := string(byteData)
	fmt.Println("Data = ", content)

}
