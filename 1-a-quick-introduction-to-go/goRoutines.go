// you can only execute functions or anonymous functions as goroutines.
package main

import (
	"fmt"
	"time"
)

func myPrint(start, finish int) {
	for i := start; i <= finish; i++ { // 0 to 4, 1 to 4 ...
		fmt.Print(i, " ")
	}
	fmt.Println()
	time.Sleep(100 * time.Microsecond) // 100 microseconds
}

func main() {
	for i := 0; i < 1; i++ { // 0, 1, 2 ...
		go myPrint(i, 5) // will execute in random order; The Go scheduler is responsible for the execution of goroutines
	}
	time.Sleep(time.Second) // 1 sec
}
