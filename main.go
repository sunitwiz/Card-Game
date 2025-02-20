package main

import (
	"fmt"
)

func main() {

	d := newDeck() // Ensure the function call is newDeck
	//var d1, d2 deck
	d1, _ := deal2(d, 5)
	//d.print()
	d1.print()
	//d2.print()
	fmt.Println("Hello")
	//fmt.Println("d1")

	var appenddesck string
	//appenddesck = strings.Join(d1, " ")

	fmt.Println(appenddesck)
	fmt.Println([]byte(appenddesck))

	d.saveTofile("deck.txt")

	readFiledeck := NewdeckFromFile("deck.txt")
	readFiledeck.print()

}
