package main

import (
	"echo-blog/pointer"
)

func main() {
	// booksapi.Run()
	pointer.Run()
}

func Hello(name string, language string) string {
	const englishHelloPrefix = "hello, "
	if name == "" {
		name = "World"
	}
	switch language {
	case "Spanish":
		return "Hola, " + name
	case "Chinese":
		return "你好, " + name
	}
	return englishHelloPrefix + name
}
