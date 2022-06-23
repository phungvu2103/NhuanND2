package main

import (
	"fmt"

	"example.com/greetings/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("NhuanND")
	fmt.Println(message)
}
