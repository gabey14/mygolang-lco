package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	// content is in the byte format
	content, _ := ioutil.ReadAll(response.Body)
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is:", byteCount)

	fmt.Println(responseString.String())

	fmt.Println(string(content))
}
