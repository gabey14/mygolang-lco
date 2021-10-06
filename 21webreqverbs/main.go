package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video - LCO")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	checkNilErr(err)

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

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price":0,
			"platform":"learncodeonline.in"
		}
	`)
	response, err := http.Post(myUrl, "application/json", requestBody)

	checkNilErr(err) // panic if error

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {

	const myUrl = "http://localhost:8000/postform"

	// Form Data

	data := url.Values{}
	data.Add("firstname", "Abey")
	data.Add("lastname", "George")
	data.Add("email", "abeygeorge14@gmail.com")

	// Reponsiblity of the user to close the response body
	response, err := http.PostForm(myUrl, data)
	checkNilErr(err)

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	fmt.Println(string(content))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
