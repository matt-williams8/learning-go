package main

import "fmt"

const delimiter = ", "
const englishHelloPrefix = "Hello"

var languageToPrefix = map[string]string{
	"en": englishHelloPrefix,
	"fr": "Bonjour",
	"es": "Hola",
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefixForLanguage(language) + delimiter + name
}

func greetingPrefixForLanguage(language string) (prefix string) {
	prefix = languageToPrefix[language]
	if prefix == "" {
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("World", "en"))
}
