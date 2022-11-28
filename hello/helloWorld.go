package hello

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world
type Language string

var (
	None    Language = "English"
	Spanish Language = "Spanish"
	French  Language = "French"
)

const englishHellPrefix = "Hello, "
const spanishHellPrefix = "Hola, "
const frenshHellPrefix = "Bonjour, "

func Hello(name string, language Language) string {
	if name == "" {
		return englishHellPrefix + "world"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language Language) string {
	switch language {
	case Spanish:
		return spanishHellPrefix
	case French:
		return frenshHellPrefix
	default:
		return englishHellPrefix
	}
}
