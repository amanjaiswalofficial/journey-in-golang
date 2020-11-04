// SLICING EXAMPLE
package main

import (
	"fmt"
	"strings"
	// Sub-package (ref: Documentation)
	"io/ioutil"
)


type deck []string

func (d deck) print(){

	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	// Will return a deck type hence initializing the same 
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


func deal(d deck, handSize int) (deck, deck){ 

	return d[:handSize], d[handSize:]
}


/*
	TYPE CONVERSION: To convert value of one type into anothe
	Ex - []string(<VARIABLE>), will convert a <VARIABLE> into
	a slice of string
	Similar, a slice of byte or byte-slice, to convert a string into byteslice
	i.e. []byte
*/
func (d deck) toString() string {

	// Method to join values from a slice of string, using a separator
	return strings.Join([]string(d), ",")
}

/*
	WRITING TO A FILE
	Here deck is passed as receiver
	Whereas filename as argument
	0666 refers to permission: anyone can read or write on this
	From documentation, we know this method returns an error if exists
*/
func (d deck) saveToFile(fileName string) error {
	// First []byte i.e. byteslice conversion of deck's to String
	// Then write the file with fileName
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}