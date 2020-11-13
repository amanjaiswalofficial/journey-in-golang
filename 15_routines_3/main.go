/*
	This example deals with
	1. Running infinite loops, in 2 ways
	2. Writing literal in go
	3. Using time.sleep to sleep for x seconds
	4. Possible bug in sending variable by reference to literal
	5. Passing variable by value to a literal in go
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.amazon.in",
		"http://www.google.com",
		"http://www.facebook.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checklinkUsingChannel(link, c)
	}

	/* 
		1. Running this in an infinite loop
	*/
	// for {
	// 	go checklinkUsingChannel(<-c, c)
	// }

	/*
		2. Alternative way for the same,
		this time all the blocking calls are stacked once
		then they are looped one by one
	*/
	// for l := range c {
	// 	go checklinkUsingChannel(l, c)
	// }

	/*
		The objective of method is to check links after some time
		To enable this Time.sleep() can be used
		But to make sure that it doesn't stop the main routine's flow
		The time.sleep() should be called inside a function and not
		in main routine itself.

		For, this, literal in go (lambda in python) are used
		Ex - 
		func(){

		}()
	*/

	/*
		In this case, a warning of l caught in literal
		Due to the fact that once all 3 links are executed
		The func still reads the value as main method
		Hence, it gets stuck on the last value and keeps executing
		The loop for the same
	*/
	// for l := range c {
	// 	go func() {
	// 		time.Sleep(5 * time.Second)
	// 		checklinkUsingChannel(l, c)
	// 	}()
	// }

	/*
		As a fix to this
		The value of l, can be passed on as a value
		Instead of the same by passing the reference
		Here, its passed as a string to the method

		Due to which, now a copy of original value is sent
		And not the original value
		So, it can keep on changing every 5 seconds

	*/
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checklinkUsingChannel(link, c)
		}(l)
	}
	/* 
		As expected
		the above method runs every 5 seconds
		And checks for status of the urls
	*/

}

func checklinkUsingChannel(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error opening: ", link)
		c <- link
		return
	}

	fmt.Println(link, "is running")
	c <- link
}
