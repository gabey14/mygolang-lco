package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Walrus syntax
	welcome := "Welcome to user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for our Pizza: ")

	// comma ok syntax || err err syntax
	input, _ := reader.ReadString('\n')

	fmt.Print("Thanks for rating - ", input)
	fmt.Printf("Type of this rating is %T ", input)

}
