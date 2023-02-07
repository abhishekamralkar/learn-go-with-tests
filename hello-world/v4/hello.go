package main

import "fmt"

const (
	englishHelloPrefix = "Hello "
	hindinamastePrefix = "Namaste "
	spanishholaPrefix  = "Hola "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "hindi" {
		return hindinamastePrefix + name
	}
	if language == "spanish" {
		return spanishholaPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Anay", "hindi"))
}
