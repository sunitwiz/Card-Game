package main

import (
	"fmt"
)

func main() {

	d := newDeck() // Ensure the function call is newDeck

	d1, _ := deal(d, 5)
	//d.print()
	d1.print()
	//d2.print()
	fmt.Println("Hello")
	//fmt.Println("d1")

	var appenddesck string
	//appenddesck = strings.Join(d1, " ")

	fmt.Println(appenddesck)
	fmt.Println([]byte(appenddesck))

	d.saveToFile("deck.txt")

	readFiledeck := newDeckFromFile("deck.txt")
	readFiledeck.print()
	d.shffle()
	d.print()

}
