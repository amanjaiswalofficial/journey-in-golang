// This example deals with
// 1. Type conversion
// 2. Byte-Slice or Slice of Bytes
// 3. Writing to a file
// 4. Reading from a file
// 5. Creating a random number everytime using a new seed value for it
// 6. Usage of index in range/loop, Using len method
// Creation: toString(), saveToFile(), readFromFile(), shuffle()

package main

import "fmt"

func main() {

	cards := newDeck()
	
	fmt.Println(cards.toString())

	// Passing fileName to save as
	cards.saveToFile("my_file")

	fmt.Println("-------")
	fmt.Println("File Saved Containing deck")
	fmt.Println("-------")
	fmt.Println("Reading data from file")
	fmt.Println("-------")

	cardsRead := newDeckFromFile("my_file")
	fmt.Println(cardsRead.toString())

	fmt.Println("-------")
	fmt.Println("Shuffling Cards")
	fmt.Println("-------")

	cardsRead.shuffle()
	fmt.Println(cardsRead.toString())

}