package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"A": 1,
		"B": 2,
	}

	if number, isExist := m["A"]; isExist == true {
		fmt.Println(number)
	}
	//
	number := 90
	guess := 95
	if guess < 10 {
		fmt.Println("Less than 10")
	}
	if guess > 100 {
		fmt.Println("Greater than 100")
	}
	if guess < number {
		fmt.Println("Less than exist number")
	}
	if guess > number {
		fmt.Println("Greater than exist number")
	}

	// || vs &&
	if guess < 10 || guess > 100 {
		fmt.Println("Less than 10 or Greater than 100")
	}

	//switch case
	exam := 90
	switch exam {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("Default")
	}
}
