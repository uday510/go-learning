package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if (len(d)) != 16 {
		t.Errorf("Expected deck length of 16, but got, %v", len(d))
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

_:
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := deck.readFromFile("cards.txt")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}
}
