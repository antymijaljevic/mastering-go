// log.Fatal() function is used when something erroneous has happened and you just want to exit your program
// log.Panic() implies that something really unexpected and unknown

package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		log.Fatal("Fatal: Hello World!") // exit status 1
	}
	log.Panic("Panic: Hello World!") // gives memory addresses (pointers)
}
