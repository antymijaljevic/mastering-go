package main

import (
	"fmt"
	"os"
	"path"
)

// Data type for holding the records of the phone book that is a Go structure with three fields
// Structures group a set of values into a single data type
type Entry struct {
	Name    string
	Surname string
	Tel     string
}

// There exists a global variable that holds the data of the phone book, which is a slice of structures named data
var data = []Entry{}

// The search() function performs a linear search on the data slice. Linear search is slow, but it does the job for now considering that the phone book does not contain lots of entries.
func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
			return &data[i] // The use of &data[i] instead of data[i] in your Go code is a matter of returning a reference (pointer) to the data rather than a copy of the data itself
		}
	}
	return nil
}

// The list() function just prints the contents of the data slice using a for loop with range
func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		exe := path.Base(arguments[0]) // The exe variable holds the path to the executable file
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	// append data
	data = append(data, Entry{"Mihalis", "Tsoukalos", "2109416471"})
	data = append(data, Entry{"Mary", "Doe", "2109416871"})
	data = append(data, Entry{"John", "Black", "2109416123"})

	// Differentiate between the commands
	switch arguments[1] {

	// user search
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}

		result := search(arguments[2])
		if result == nil {
			fmt.Println("Entry not found:", arguments[2])
			return
		}
		fmt.Println(*result)
	// list persons
	case "list":
		list()
	default:
		fmt.Println("Not a vaild option")
	}
}
