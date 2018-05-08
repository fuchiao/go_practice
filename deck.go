package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardNumbers := []string{"Ace", "Jack", "Queen", "King"}
	var d deck

	for _, suit := range cardSuits {
		for _, num := range cardNumbers {
			d = append(d, num+" "+suit)
		}
	}
	return d
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) {
	err := ioutil.WriteFile(filename, []byte(d.toString()), 0644)
	if err != nil {
		fmt.Println("Error: failed to write file", filename)
	}
}

func NewDeckFromFile(filename string) deck {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: failed to read file", filename)
	}
	s := strings.Split(string(content), ",")
	return deck(s)
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
