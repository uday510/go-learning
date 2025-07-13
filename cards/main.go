package main

import (
	"fmt"
	"time"
)

var buildTime string
var gitCommit string
var goVersion string

func main() {
	fmt.Println("Compiled at:", buildTime)
	fmt.Println("Git commit: ", gitCommit)
	fmt.Println("Go version: ", goVersion)

	var card = "Ace of Spades"

	fmt.Println("Running at:", time.Now().UTC().Format("2006-01-02 15:04:05 UTC"))
	fmt.Println(card)

}
