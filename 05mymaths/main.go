package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Hello to maths in golang")

	// var mynumberOne int = 2
	// var mynumberTwo float64 = 4

	// fmt.Println("The sum is: ", mynumberOne+int(mynumberTwo))

	// random number / math/rand package
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// random number  from / cryto/rand package
	// Much more reliable because of crypto 
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))

	fmt.Println(myRandomNum)

	
}
