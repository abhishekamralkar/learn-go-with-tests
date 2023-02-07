package main

import "fmt"

const (
	englishHelloPrefix  = "Hello "
	frenchBonjourPrefix = "Bonjour "
	spanishHolaPrefix   = "Hola "
	hindiNamastePrefix  = "Namaste "

	spanish = "Spanish"
	hindi   = "Hindi"
	french  = "French"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchBonjourPrefix
	case spanish:
		prefix = spanishHolaPrefix
	case hindi:
		prefix = hindiNamastePrefix
	default:
		prefix = englishHelloPrefix
	}

	return

}

func main() {
	fmt.Println(Hello("Anay", "Hindi"))
}
