package main

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
