package main

import "fmt"

type bot interface {
	getGreeting() string
	test() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	// sb := spanishBot{}

	printGreeting(eb)
	// printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func getGreeting() string {
	return "Tes"
}

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (englishBot) test() string {
	return "Test"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
