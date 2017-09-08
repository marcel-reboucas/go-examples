package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("The length of the new deck must be 52, but got %v", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("The first card should be Ace of Hearts, but got %s", d[0])
	}

	if d[len(d)-1] != "King of Spades" {
		t.Errorf("The last card should be King of Spades, but got %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "deck_testing"
	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)
	newDeck := newDeckFromFile(filename)

	if len(newDeck) != len(d) {
		t.Errorf("Expected %v cards in the deck, got %v instead", len(d), len(newDeck))
	}

	os.Remove(filename)
}
