package main

import (
	"encoding/json"
	"fmt"
)

// private when first letter is lowercase
// public when first letter is uppercase
type course struct {
	Name      string   `json:"coursename"`
	Price     int      `json:"price"`
	Platform  string   `json:"website"`
	Passsword string   `json:"-"`
	Tags      []string `json:"tags,omitempty"` // slice of strings
}

func main() {
	fmt.Println("Welcome to json video")
	EncodeJson()

}

func EncodeJson() {
	lcoCourses := []course{
		{"React Js Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}
