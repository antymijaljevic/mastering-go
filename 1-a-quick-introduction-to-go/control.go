package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a command line argument")
		return
	}
	argument := os.Args[1] // go run control.go something
	// fmt.Println(argument)

	// With expression after switch
	switch argument {
	case "0":
		fmt.Println("Zero!")
	case "1":
		fmt.Println("One!")
	case "2", "3", "4":
		fmt.Println("2, 3 or 4")
		fallthrough // The fallthrough keyword tells Go that after this branch is executed, it will continue with the next branch
	default:
		fmt.Println("Value", argument)
	}

	printType(argument) // string type

	// go doc strconv.Atoi
	value, err := strconv.Atoi(argument) // converts string to int A-to-i
	if err != nil {                      // if you return some error
		fmt.Println("Cannot convert to int:", argument)
		return
	}
	printType(value) // int type

	// No expression after switch
	switch {
	case value == 0:
		fmt.Println("Zero")
	case value > 0:
		fmt.Println("Positive integer")
	case value < 0:
		fmt.Println("Negative integer")
	default:
		fmt.Println("This should not happen:", value)
	}
}

// Check type of variable (checking type isn't covered yet)
func printType(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Printf("Type of %v: %v\n", v, t)
}
