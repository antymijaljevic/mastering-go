package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args     // everything is considered a string in os.Args as os.Args is a sclice of strings
	if len(arguments) == 1 { // The first command-line argument stored in the os.Args slice is always the name of the executable
		fmt.Println("Need one or more arguments!")
		return // exit main function
	}

	var min, max float64

	for i := 1; i < len(arguments); i++ { // skip arg 0 (executable name)
		n, err := strconv.ParseFloat(arguments[i], 64) // convert args from string to float64
		if err != nil {                                // if returns error start from the beginning of the loop
			// fmt.Println("can't convert .. skipping this arg")
			continue
		}

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max", max)
}
