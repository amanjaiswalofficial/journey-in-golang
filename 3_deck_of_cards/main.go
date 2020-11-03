// This example deals with
// 1.Using custom types and receiver methods
// Creations: deck.go -> deck, print()

package main

func main() {

	// deck since part of same package (main), can be used to refer
	// to this array of strings
	cards := deck{"Ace of Diamonds", "Ace of Spades"}
	cards = append(cards, "Ace of Hearts")

	// Calling print of deck from a deck type variable
	cards.print()
}