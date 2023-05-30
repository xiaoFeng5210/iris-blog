package main

import (
	"echo-blog/pointer"
)

func main() {
	// booksapi.Run()
	pointer.Run()
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		return "Hola, "
	case "Chinese":
		return "你好, "
	default:
		return "hello, "
	}
}
