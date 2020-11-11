/*
This example deals with
	1. Using Writer interface
	2. Writing your own Writer interface
	3. Using io.Copy to copy data from source to destination
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	/*
		An alternative to the previous approach
		A Writer can be used to write something to a destination, Ex- stdout/file etc
		Here,
		io.Copy takes 2 args -> 1 that implements Writer interface
		Another that implements Reader interface (Visit documentation for more)

		This way, on calling copy for them
		The data from body of response, is written to stdout, hence printed on screen
		On running
		'go run main.go'
	*/
	// UNCOMMENT TO RUN
	// io.Copy(os.Stdout, resp.Body)



	/*
		Using your custom types for the same purpose
		Here, a custom struct logWriter is used
		logWriter implements method Write, which
		: takes a byteslice as input
		: returns an int and an error as output

		Both properties match to Write method of a Writer interface
		Hence, in place of os.Stdout, we can use logWriter

		As both are now Writers
		Because, both implement Write method of same kind as required
	*/
	lw := logWriter{}

	// Hence, running with io.Copy(lw...)
	io.Copy(lw, resp.Body)

}

/*
	IMPORTANT: Even though the Write method has similar contract as required
	For any Writer interface type

	It doesn't guarentee that it will operate the same way
	Which is dependent completely on how the method is written

	Write method implementation only makes it qualified to be an input
	for io.Copy
	Whether or not, io.Copy will run properly on it depends on method def.
*/
func (logWriter) Write(bs [] byte) (int, error) {

	// different logic than Writer's Write method but still fulfills
	// the requirement, i.e. to print the body
	fmt.Println(string(bs))
	fmt.Println("Just wrote bytes", len(bs))
	return len(bs), nil
}