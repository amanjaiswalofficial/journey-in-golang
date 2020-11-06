/*
	This example deals with
	1. Writing tests
	2. Running solo tests and all tests
	3. Displaying error on failure
	4. Deleting a file using os.remove()
	Creations: TestNewDeck(), TestSaveToDeckAndNewDeckFromFile()
*/

package main

import (
	"os"
	"testing"
)

/*
	To simply run all the tests, use
	'go test',
	Which returns as OK or FAIL depending if all tests passed or failed
	To run a specific test, run test from VS Code can be used
	OR
	'go test deck.go deck_test.go',
	Because newDeck is used, deck.go is required to be used
*/
func TestNewDeck(t *testing.T) {
	d := newDeck()

	// if len doesn't match, give error
	if len(d) != 9 {
		t.Errorf("Expected length to be 16 but got %v", len(d))
	}

	// if first card isn't Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("First card isn't Ace of Spades, it is: %v", d[0])
	}

	// if last card isn't ...
	if d[len(d)-1] != "Three of Hearts" {
		t.Errorf("Last card isn't Three of Hearts, it is %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	// As part of cleanup, first we delete if any file that was created last time
	// That couldn't be deleted because the program crashed in t.errorF
	os.Remove("_decktesting")

	deck := newDeck()
	// Save to a file with a name
	deck.saveToFile("_decktesting")

	// Load the same file
	loadedDeck := newDeckFromFile("_decktesting")

	// if content doesn't match, then raise error
	if len(loadedDeck) != 9 {
		t.Errorf("Expected 9 cards, got %v", len(loadedDeck))
	}

	// if t.Errorf didn't execute, meaning test passed, so delete file
	os.Remove("_decktesting")
}
