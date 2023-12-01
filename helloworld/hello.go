package main

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const igboHelloPrefix = "Kedu, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	case "Igbo":
		prefix = igboHelloPrefix
	}

	return prefix + name
}
func main() {

	// fmt.Println(Hello("Chris"))
}
