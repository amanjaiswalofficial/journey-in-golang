package main

import "fmt"

/* 1. Declaring a variable
func main() {
	// Long form, explicitly defining the type for variable
	//var card string = "Ace of Spades"

	// Another way, relies on Go to identify the type for the variable
	// NOTE: Only using := for first time definition, afterwards = for assignment
	card := "Ace of Spades"
	fmt.Println(card)

}
*/

/* 2. Using a method to return a value with type
func main() {
	card := newCard()
	fmt.Println(card)

}

// string helps Go identify what kinda data it is supposed to return
func newCard() string {
	return "Ace of Spades"
}
*/

// 3. Using arrays
// NOTES: Arrays or Slices: Arrays are of fixed value/size, once declared don't change
// Whereas, Slices have usual list based operations
// Due to this, later the appended value in cards below is saved into a new instance of cards
func main() {
	// defining the type of values for this, i.e. array of strings
	cards := []string{"Ace of Diamonds", newCard()}

	// append new value to the array
	cards = append(cards, "Ace of Hearts")

	// run a for loop, to print index and value from array
	for i, card := range cards {
		fmt.Println(i, card)
	}

}

// string helps Go identify what kinda data it is supposed to return
func newCard() string {
	return "Ace of Spades"
}
