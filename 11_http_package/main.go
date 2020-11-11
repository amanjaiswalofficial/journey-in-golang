/*
This example deals with
1. Using http package
2. Using make() to init byteslice
3. Using Body.Read() to read body of a Response object
*/
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// package used to make GET Requests
	resp, err := http.Get("http://google.com")

	/*
		If error, then exit
	*/
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	/*
		Another way to initialize empty variables
		Here, bs is a slice containing bytes
		And with defined size of 9999
	*/
	bs := make([]byte, 9999)
	// From documentation, find this method to print Body
	// After passing its content to a byteslice
	resp.Body.Read(bs)
	// Convert byteslice to a slice of strings
	// Then print
	fmt.Println(string(bs))

}
