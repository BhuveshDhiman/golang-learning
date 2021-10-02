package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gk848f"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println("URL = ", myUrl)

	// parsing
	result, _ := url.Parse(myUrl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Printf("The type of query params is: %T\n", qParams)
	fmt.Println(qParams["coursename"])
	fmt.Println(qParams["paymentid"])

	for _, val := range qParams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=bhuvesh",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
