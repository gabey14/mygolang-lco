package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	// NOTE : map[key]value
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)

	fmt.Println("JS shorts for:", languages["JS"])

	// NOTE : Delete a key value pair
	delete(languages, "RB")

	fmt.Println("List of all languages: ", languages)

	// > Loops are interesting in Golang
	// > comma ok syntax
	// for _, value := range languages {
	// 	fmt.Printf("For key v, value is %v\n", value)
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)

	}

}
