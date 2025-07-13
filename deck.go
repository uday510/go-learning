package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Printf("%2d: %s\n", i+1, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(fileName string) {
	fmt.Println("saving to the file ", fileName, "...")
	err := os.WriteFile(fileName, []byte(d.toString()), 0666)
	if err != nil {
		fmt.Println("Error while saving", err)
		os.Exit(-1)
	}
}

func (d deck) readFromFile(fileName string) deck {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error while reading deck", err)
		os.Exit(1)
	}

	return strings.Split(string(data), ",")
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for idx := range d {
		newPos := random.Intn(len(d) - 1)
		d[idx], d[newPos] = d[newPos], d[idx]
	}
}
