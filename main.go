package main

import (
	"cardgame/database"
	"fmt"
	"strings"
)

func decktest() {
	d := newDeck() // Ensure the function call is newDeck

	d1, _ := deal(d, 5)
	d.print()
	d1.print()

	fmt.Println("Hello")
	fmt.Println("d1")

	appenddesck := strings.Join(d1, " ")

	fmt.Println(appenddesck)
	fmt.Println([]byte(appenddesck))

	d.saveToFile("deck.txt")

	readFiledeck := newDeckFromFile("deck.txt")
	readFiledeck.print()
	d.shuffle()
	d.print()

}

func structtest() {

	alex := person{}
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Printf("%+v", alex)
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim_anderson@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updatecontactpointer("new_jum_anderson@gmail.com") //it won't update

	jim2 := jim.updatecontact2("new_jum_anderson2@gmail.com") //it will update

	jum3 := &jim2
	jum3.updatecontactpointer("new_jum_anderson3@gmail.com") //it will also update
	//or directly
	jim2.updatecontactpointer("new_jum_anderson4@gmail.com") //it will also update

	fmt.Printf("%+v", jim2)
	fmt.Printf("%+v", jim)

	mp := map[string]int{ //map declaration
		"one":   1,
		"zwo":   2,
		"three": 3,
	}

	fmt.Printf("%+v\n", mp)
	mp["three"] = 4

	fmt.Println("\n", mp)
	delete(mp, "zwo") //deletes an item from the map
	fmt.Println(mp)
	fmt.Println(mp["pick"])
	val, exists := mp["pick"] //checks if the key exists
	fmt.Println(exists, "the value is", val)

	for key, value := range mp { //iterates over the map
		fmt.Println(key, value)
	}

}

func main() {

	decktest()

	structtest() //structure and maps

	database.Connect()

}
