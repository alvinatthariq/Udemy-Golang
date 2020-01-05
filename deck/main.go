package main

func main(){
	cards := newDeckFromFile("my_deck")
	cards.print()
	cards.shuffle()
	cards.print()

}
