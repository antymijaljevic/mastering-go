package main

import "fmt"

func main() {
	// Get user input
	fmt.Printf("Please give me your name: ") // not adding a newline character
	var name string
	fmt.Scanln(&name) //  part is taking the address of the name variable; reads input until a newline character is encountered
	fmt.Println("Your name is", name)
}
