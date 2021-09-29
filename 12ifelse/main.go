package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("If else in golang")

	loginCount := 10
	var result string

	// NOTE : If else if else

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// NOTE : For Web
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}

	var err error = errors.New("Error")
	if err == nil {
		fmt.Println(err)
	} else {
		fmt.Println(err)
	}

}
