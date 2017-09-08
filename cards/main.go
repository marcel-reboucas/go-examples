package main

type color string

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	newDeckFromFile("my_cards")

	cards.shuffle()
	cards.print()
}
