package main

import "fmt"

// First letter Capital means its public
const LoginToken string = "dhsgfiwgewr" // Public

func main() {
	// String
	var username string = "abey"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	// Boolean
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// Small Int
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// Small Float
	var smallFloat float32 = 255.453242532523
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString)

	// Implicit Type
	var website = "learncodeonline.in"
	fmt.Println(website)

	// No var style only allowed in method not globally - Walrus operator
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}
