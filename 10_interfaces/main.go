package main

import "fmt"

/*
	An interface defined
	Which tells that each of the other types that
	Have a method with:

	name: getGreeting()
	returning: string

	Can be a bot too
	Hence, any method on bot type can be used on them as well.
*/

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	// Here, both eb and sb can call printGreeting
	// And since both have required method
	// Both are bots and hence no error in execution
	printGreeting(eb)
	printGreeting(sb)

}

func (englishBot) getGreeting() string {
	return "Hi There"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

/*
	Updated printGreeting
	This time, accepts a bot
	And as per definiton of bot
	There will be a method named, getGreeting()
	That returns a string

	Hence, both englishBot and spanishBot
	Can call this method and their getGreeting()
	method will be called.

	Helping in both reducing repeated code
	And reducing redundancy
*/
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

/*
	Notes:
	Interfaces are implicit
	We didn't have to declare that englishBot is of type bot
	It automatically understood based on the declaration of methods
	That englishBot type did.
*/