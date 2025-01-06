package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Separate your "domain" code from the outside world.
func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
