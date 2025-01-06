package main

import "fmt"

// Separate your "domain" code from the outside world.
func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
}
