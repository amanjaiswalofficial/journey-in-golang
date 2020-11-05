// SLICING EXAMPLE
package main

import (
	"fmt"
	"os"
	"strings"
	// Sub-package (ref: Documentation)
	"io/ioutil"
	"time"
	"math/rand"
)

type deck []string

func (d deck) print() {

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

func deal(d deck, handSize int) (deck, deck) {

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

/*
	READING A FILE using ReadFile
	Returns 2 values, a byteSlice and error
	if some error existed, exit the program
	Otherwise, do type conversion from byteslice to String
	Then Split the string using sep=","
	ANd return a deck converted value for this
*/
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	// Returning a deck as expected
	return deck(s)

}

/*
	SHUFFLE deck of cards
	IMPORTANT: Since random function to generate random number uses a seed
	And same seed value is used everytime rand is used
	To ensure every time a different seed value is used
	A new source is required every time
	For a new source to be used every time, time.UnixNano is used
	Which returns an int value which returns a randomInt every time
	This makes a new source everytime which makes a new seed for rand
*/
func (d deck) shuffle(){
	randomInt := time.Now().UnixNano()
	source := rand.NewSource(randomInt)
	r := rand.New(source)

	// Shuffle cards by switching indexes of values
	for i := range d{
		// Generate new random values every time
		newPosition := r.Intn(len(d) - 1)

		d[i],d[newPosition] = d[newPosition],d[i]

	}

	// Returning nothing as only shuffling existing is required
}