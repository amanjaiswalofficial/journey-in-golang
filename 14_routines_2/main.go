/*
	This example deals with
	1. Creating channels
	2. Using channel to communicate between routines
	3. Fixing the bug caused while using routines without channels
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

	/*
		A channel is a communication medium as name suggests
		Made by make() it also expects a type
		This type will be used to communicate data b/w routines
		i.e. here, only strings can be communicated via channel c
	*/
	c := make(chan string)

	for _, link := range links {
		// passing the channel into modified checklink
		// now accepting a channel
		go checklink(link, c)
	}

	/*
		The receiver 
		i.e. fmt.Println(<-c)
		[for msg coming from a channel]
		Is a blocking call, which only when executed
		Lets the code move forward
		Hence, for as many links as used
		It must block the code for same amount of times
		To make sure the status for each link is printed
	*/
	for i:=0; i<len(links);i++{
		// for 3 links, it will make 3 blocking calls
		// And wait 3 times. then will move the rest of code
		// <- c to receive message from channel and print it
		fmt.Println(<-c)
	}

}

func checklink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error opening: ", link)
		// c <- "" to send message to the channel, only string as defined
		c <- "Might be down"
		return
	}

	fmt.Println(link, "is running")
	c <- "Working successfully"
}
