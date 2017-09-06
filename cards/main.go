package main

type color string

func main() {
	cards := newDeck()

	hand, rest := deal(cards, 5)
	hand.print()
	rest.print()
}
