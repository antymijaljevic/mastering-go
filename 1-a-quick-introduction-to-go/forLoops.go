package main

import "fmt"

func main() {
	// Traditional for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}

	// idiomatic Go: conventions and style guidelines established by the Go programming language community
	i := 0
	for ok := true; ok; ok = (i != 10) { // ok is true; do until ok true; ok is true so continue the loop until ok false
		fmt.Println(i*i, " ") // the square
		i++
	}
	fmt.Println()

	// For loop used as while loop
	x := 0
	for {
		if x == 10 {
			break
		}
		fmt.Print(x*x, " ")
		x++
	}
	fmt.Println()

	// this is a slice but range also works with arrays
	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range aSlice { // use _ in the place of the value that you want to ignore
		fmt.Println("index", i, "value", v)
	}
}
