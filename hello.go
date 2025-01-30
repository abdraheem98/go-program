package main

import "fmt"

func main() {
	age := 18
	switch {
	case age < 18:
		fmt.Println("you'r a minor")
	case age >= 18 && age < 60:
		fmt.println("you're an adult")
	default:
		fmt.println("you're an sz")
	}
}
