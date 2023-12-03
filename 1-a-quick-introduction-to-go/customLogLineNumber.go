package main

import "fmt"

func Print[T any](s []T) { // Generics allow you to write functions and data types that are "generic" over some type(s), without specifying the exact types they operate on
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	Ints := []int{1, 2, 3}
	Strings := []string{"One", "Two", "Three"}
	Print(Ints)
	Print(Strings)
}
