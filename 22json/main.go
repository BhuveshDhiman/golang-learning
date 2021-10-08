package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"courseName"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON in Go")
	// EncodeJson()
	decodeJson()

}

func EncodeJson() {
	courses := []course{
		{"ReactJS", 299, "AWS", "12345678react", []string{"web-dev", "js"}},
		{"MERN", 199, "AWS", "12345678mern", []string{"full-stack", "js"}},
		{"React-Native", 499, "AWS", "12345678rn", nil},
	}

	// package as JSON data
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func decodeJson() {
	jsonData := []byte(`
	{
		"courseName": "ReactJS",
		"price": 299,
		"platform": "AWS",
		"tags": ["web-dev","js"]
	}
	`)

	var courseData course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonData, &courseData)
		fmt.Printf("%#v\n", courseData)
	} else {
		fmt.Println("JSON was not valid")
	}

	// just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}

}
