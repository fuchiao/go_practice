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
