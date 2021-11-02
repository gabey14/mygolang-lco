package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gabey14/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")

	fmt.Println("Listening at port 4000...")
	log.Fatal(http.ListenAndServe(":4000", r))

}
