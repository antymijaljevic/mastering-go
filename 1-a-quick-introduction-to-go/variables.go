package main

import (
	"fmt"
	"math"
)

var Global int = 12345
var AnotherGlobal = -5678

func main() {
	var j int // 0
	i := Global + AnotherGlobal
	fmt.Println("Initial j value", j)
	j = Global
	// math.Abs() requires a float64 parameter so we type cast it appropriately
	k := math.Abs(float64(AnotherGlobal)) // Explicit conversion from int to float64
	fmt.Printf("Global=%d, i=%d, j=%d k=%.2f.\n", Global, i, j, k)
}
