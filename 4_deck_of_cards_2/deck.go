package main

import "fmt"

type deck []string

func (d deck) print(){

	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	// Will return a deck type hence initializing the same 
	// and write same in declaring the function
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}


// Takes 2 input, one of type deck, to be used as d
// Another of type int, to be used as handSize
// Returns 2 values, both type deck, hence can use print() on them
func deal(d deck, handSize int) (deck, deck){ 

	// SLICING EXAMPLE, to return 2 deck type values
	return d[:handSize], d[handSize:]
}