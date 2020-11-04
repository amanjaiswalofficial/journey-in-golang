// This example deals with
// 1. Looping over data
// 2. Create a method to return data of deck
// 3. Create a method to return SLICES of a deck, both of type deck
// Creation: newDeck(), deal()

package main

import "fmt"

func main() {

	cards := newDeck()
	hand, remainingCards := deal(cards, 5)

	fmt.Println("The hand is:")
	hand.print()

	fmt.Println("The remaining cards are:")
	remainingCards.print()
}