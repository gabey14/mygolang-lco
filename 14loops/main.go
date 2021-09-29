package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// NOTE : Type of days
	// fmt.Printf("Type of days: %T\n", days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// NOTE : This will give the index
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	// for _, day := range days {
	// 	fmt.Printf("value is %v\n", day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		// NOTE  : goto is used to jump to a label
		if rougueValue == 2 {
			goto lco
		}
		// if rougueValue == 4 {
		// break
		// }

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

	// NOTE : This is a label
lco:
	fmt.Println("Jumping at learncodeonline.com")

}
