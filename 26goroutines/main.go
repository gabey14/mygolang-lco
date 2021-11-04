package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// NOTE - Go Routines Quote
// > Do not communicate by sharing memory;instead share memory by communicating.

var signals = []string{"test"}

// Wait group variable
// Usually they are pointers
var wg sync.WaitGroup // pointer

var mut sync.Mutex //pointer

func main() {
	// go greeter("Hello")
	// greeter("world")

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		// wg.Add(2)
		wg.Add(1)
	}
	// for _, web := range websitelist {
	// 	go getStatusCode(web)
	// }

	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {
	// passes a signal done
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		// fmt.Println("OOPS in endpoint")
		log.Fatal(err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}

// func greeter(s string) {

// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}

// }
