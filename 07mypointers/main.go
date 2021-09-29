package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// NOTE : Declare a pointer
	// var ptr *int

	// NOTE : Default value of a pointer is nil
	// fmt.Println(ptr)

	myNumber := 23
	// NOTE : Assign a value to a pointer
	// >  Reference - &
	var ptr = &myNumber

	fmt.Print("The value of x is : ", *ptr, "\n")
	fmt.Print("The value of x is stored at : ", &ptr)

	// NOTE : Change the value of the pointer
	// NOTE : Remember that we are updating the value at the address stored in the pointer so whichever variable is pointing to that address will have that value.
	*ptr = *ptr + 2
	fmt.Println("\nThe new value of x is :", myNumber)

}
