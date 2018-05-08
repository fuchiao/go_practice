package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLen := 16

	if len(d) != expectedLen {
		t.Errorf("Expected deck length is %d, but got %d", expectedLen, len(d))
	}
}
