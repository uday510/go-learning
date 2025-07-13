package main

import (
	"fmt"
	"go-learning/logbox"
	"time"
)

var buildTime string
var gitCommit string
var goVersion string

func initialize() {
	lines := []string{
		fmt.Sprintf("Compiled at: %s", buildTime),
		fmt.Sprintf("Git commit:  %s", gitCommit),
		fmt.Sprintf("Go version:  %s", goVersion),
		fmt.Sprintf("Running at:  %s", time.Now().UTC().Format("2006-01-02 15:04:05 UTC")),
	}
	logbox.PrintBox(lines)
}

func main() {
	initialize()

	cardsOps()
}

func cardsOps() {
	cards := newDeck()
	FileName := "cards.txt"

	cards.shuffle()
	cards.saveToFile(FileName)

	cards.readFromFile(FileName).print()
}
