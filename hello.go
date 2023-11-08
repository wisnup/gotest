package main

import "fmt"

const (
	frenchLang  = "French"
	spanishLang = "Spanish"

	englishPrefix = "Hello, "
	frenchPrefix  = "Bonjour, "
	spanishPrefix = "Hola, "
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return getPrefix(lang) + name
}

func getPrefix(lang string) (prefix string) {
	switch lang {
	case frenchLang:
		prefix = frenchPrefix
	case spanishLang:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
