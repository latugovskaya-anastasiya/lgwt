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

	if lng == spanish {
		return spanishHelloPrefix + name
	}

	if lng == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}
