package main

import (
	"echo-blog/pointer"
)

func main() {
	// booksapi.Run()
	pointer.Run()
}

func Hello(name string) string {
	const englishHelloPrefix = "hello, "
	return englishHelloPrefix + name
}
