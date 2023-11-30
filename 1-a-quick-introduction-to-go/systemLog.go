package main

import (
	"log"
	"log/syslog"
)

func main() {

	sysLog, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.go") // log level SYSTEM

	if err != nil {
		log.Println(err)
		return
	} else {
		log.SetOutput(sysLog) // outputs log to /var/log/syslog or journalctl -n 10
		log.Print("Everything is fine")
	}
}

// Nov 30 21:52:19 ante-carbon-x1 systemLog.go[5535]: 2023/11/30 21:52:19 Everything is fine
// Nov 30 21:52:24 ante-carbon-x1 systemLog.go[5669]: 2023/11/30 21:52:24 Everything is fine
