// This example deals with
// 1. Type conversion
// 2. Byte-Slice or Slice of Bytes
// 3. Writing to a file
// Creation: toString()

package main

import "fmt"

func main() {

	cards := newDeck()
	
	fmt.Println(cards.toString())

	// Passing fileName to save as
	cards.saveToFile("my_file")

}