/*
	This example deals with
	1. Using routines in go
	2. The bug caused in go routines
*/

package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.amazon.in",
		"http://www.google.com",
		"http://www.facebook.com",
	}

	fmt.Println("Running without go routines")
	fmt.Println("--------")
	for _, link := range links {
		// for each link run this
		checklink(link)
	}

	fmt.Println("Running with go routines")
	fmt.Println("--------")
	for _, link := range links {
		go checklink(link)
	}
	/*
		Even though now the method checklink is being run as a go routine
		Since by default the child routines don't communicate with main routine
		The main routine simply exits without waiting for child routines
		to finish.
		To fix this, channels are used to make communication.
	*/
	fmt.Println("Exited instantly before any execution")
}

func checklink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error opening: ", link)
		return
	}

	fmt.Println(link, "is running")
}
