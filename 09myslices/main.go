package main

import (
	// NOTE : Package name can be given an alias

	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")
	// Declare a slice when we dont mention the size it is a slice
	var fruitList = []string{"Apple", "Tomato", "Peach"}

	fmt.Printf("Type of fruitlist is %T\n", fruitList)
	fmt.Println(fruitList)

	// NOTE : Append is a function that adds a new element to the end of the slice
	// var fruitLists = append(fruitList, "Mango", "Banana")

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println(fruitList)
	// fmt.Println(fruitLists)

	fruitList = append(fruitList[1:3], "Grapes")
	fmt.Println(fruitList)

	// NOTE : Interesting stuff
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 7867

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)

	// Sorting is available in slice not in array
	sort.Ints(highScores)

	fmt.Println(highScores)

	// NOTE : Boolean value

	fmt.Println(sort.IntsAreSorted(highScores))

	// NOTE : REMOVE A VALUE FROM SLICE BASED ON INDEX

	courses := []string{"reactjs", "javascript", "swift", "python", "ruby"}

	fmt.Println(courses)

	// courses = courses[:2]

	index := 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
