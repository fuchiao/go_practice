package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLen := 16

	if len(d) != expectedLen {
		t.Errorf("Expected deck length is %d, but got %d", expectedLen, len(d))
	}
}

func TestSaveDeck(t *testing.T) {
	testFile := "_decktesting"
	os.Remove(testFile)
	deck := newDeck()
	deck.saveToFile(testFile)
	_, err := os.Stat(testFile)
	if os.IsNotExist(err) {
		t.Errorf("%s should exists", testFile)
	}
}

func TestNewDeckFromFile(t *testing.T) {

}
