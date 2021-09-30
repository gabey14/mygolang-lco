package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type of: %T\n", response)

	defer response.Body.Close() //caller's responsibility to close the connection

	if response.StatusCode == 200 {
		databytes, err := ioutil.ReadAll(response.Body)

		if err != nil {
			panic(err)
		}

		content := string(databytes)

		fmt.Println(content)
	}

}
