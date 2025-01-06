package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Separate your "domain" code from the outside world.
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
