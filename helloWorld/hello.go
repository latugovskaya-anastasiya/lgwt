package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const french = "French"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "sivuple, "

func main() {
	fmt.Println(Hello("Chris", "Spanish"))
}

func Hello(name, lng string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lng) + name
}

func greetingPrefix(lng string) (prefix string) {
	switch lng {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
