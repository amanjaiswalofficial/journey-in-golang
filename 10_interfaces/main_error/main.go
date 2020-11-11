/*
	IMPORTANT: This won't run
	CALL:
		'go run main.go'

	This example deals with
	1.Creating interfaces
	2.Importance of interfaces
*/

package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {

}

/*
	Since we don't need to use the variable of englishBot/spanishBot
	No need to declare it in function's beginning
	i.e. instead of (eb englishBot), (englishBot) is fine
*/
func (englishBot) getGreeting() string {
	return "Hi There"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

/*
	Below, for both the spanish & englishBot, there's common method
	called printGreeting, with similar functionality
	Only difference in the type of argument it accepts
	In 1 case, englishBot, in another, spanishBot

	Problems, in this approach
	Lack of reusability of similar type code (calling getGreeting())
	Ambiguity of methods

	This can be fixed by using an interface
	Which will have a superior type
	And all others which belong to this type
	Can call same method with different args
	As long as all those args satisfy some condition
	Written in interface definition
*/
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// Wont run, error
func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
