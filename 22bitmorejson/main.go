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
	// EncodeJson()
	DecodeJson()

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

func DecodeJson() {

	jsonDataFromWeb := []byte(`
		{
                "coursename": "React Js Bootcamp",
                "price": 299,
                "website": "LearnCodeOnline.in",
                "tags": [ "web-dev","js"]
        }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// Some cases where you want to add data to key value pair

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for key, val := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is %T\n", key, val, val)
	}
}
