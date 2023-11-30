package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")                             // /tmp/mGo.log
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // O_WRONLY should be opened in write-only mode.
	// The call to os.OpenFile() creates the log file for writing,
	// if it does not already exist, or opens it for writing
	// by appending new data at the end of it (os.O_APPEND)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close() // f.Close() is going to be executed just before main() returns

	iLog := log.New(f, "iLog", log.LstdFlags) // log.LstdFlags representing the log message formatting & date and time in the default format
	iLog.Println("Hello there!")
	iLog.Println("Mastering Go 3rd edition!") // iLog2023/11/30 22:27:37 Mastering Go 3rd edition!
}
