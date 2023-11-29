package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Not enough arguments")
		return
	}

	var total, nInts, nFloats int
	invalid := make([]string, 0)      // keeping all non-numeric values
	for _, k := range arguments[1:] { // skip first argument
		//is it an integer?
		_, err := strconv.Atoi(k) // we just care about error to make a count
		if err == nil {
			total++
			nInts++
			continue
		}

		// is it a float
		_, err = strconv.ParseFloat(k, 64) // float64
		if err == nil {
			total++
			nFloats++
			continue
		}

		// then it is invalid
		invalid = append(invalid, k) // now we're using actual value
	}

	// print total
	fmt.Println("#read:", total, "#ints:", nInts, "#floats:", nFloats)
	if len(invalid) > total {
		fmt.Println("Too much invalid input:", len(invalid))
		// show all invalids
		for _, s := range invalid {
			fmt.Println(s)
		}
	}
}
