package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var arguments []string = os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}

	// var file string = arguments[1]      // binary name
	var path string = os.Getenv("PATH")   // echo $PATH
	pathSplit := filepath.SplitList(path) // portable way of separating a list of paths ":"

	// search for multiple binaries
	for _, v := range arguments[1:] {
		for _, directory := range pathSplit {
			fullPath := filepath.Join(directory, v) // concatenating the different parts of a path
			// Does it exist?
			fileInfo, err := os.Stat(fullPath)
			if err == nil {
				mode := fileInfo.Mode()
				// Is it a regular file?
				if mode.IsRegular() {
					// Is it executable?
					if mode&0111 != 0 { // unix permissions; mode & 0111 performs a bitwise AND operation between mode and 0111
						fmt.Println(fullPath)
						// check if there is more than one path for the same binary
						continue
					}
				}
			}
		}
	}
}
