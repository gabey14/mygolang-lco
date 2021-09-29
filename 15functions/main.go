package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	// var name string
	// fmt.Print("Enter your name: ")
	// fmt.Scan(&name)
	greeter("Abey")

	result := adder(3, 5)
	fmt.Println("Result is :", result)

	proRes, myMessage := proAdder(2, 5, 8, 7, 8, 9)
	fmt.Println("Pro result is :", proRes)
	fmt.Println("Pro message is :", myMessage)

	// Anonymous function - no name
	func(name string) {
		fmt.Println("Hello", name)
	}("Abey")

	// Closure
	func() {
		fmt.Println("Hello closure")
	}()

	// IIFE - Immediately Invoked Function Expression
	squareOf2 := func() int {
		return 2 * 2
	}()
	fmt.Println(squareOf2)

}

// Function Signature - return type and named return value
func adder(valOne int, valTwo int) (sum int) {
	sum = valOne + valTwo
	return sum
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi Pro result function"
}

func greeter(name string) {
	fmt.Printf("Namaste %v from golang\n", name)
}
