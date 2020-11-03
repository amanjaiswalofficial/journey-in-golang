// Custom types usage
package main

import "fmt"

// Create a new type 'deck', which is basically an array of strings
type deck []string

/*
Receivers in Go:
Here print method accepts a receiver for it
i.e. a variable of type deck will be used to call print method
so d deck before print tells that a variable of type deck
Which can be referred by name d inside the method iteslf
*/
func (d deck) print(){

	// Simply iterating over d, like before done with cards
	for i, card := range d {
		fmt.Println(i, card)
	}
}