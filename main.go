package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	character string = "gitops flux series"
	sleepTime string = "1s"
)

func main() {
	for {
		// get hostname from os
		hostname, err := os.Hostname()
		if err != nil {
			log.Panicln("Error getting hostname.")
		}

		// print message
		if len(os.Getenv("CHARACTER")) != 0 {
			character = os.Getenv("CHARACTER")
		}
		fmt.Printf("hostname: %s - %s\n", hostname, character)

		// sleep time
		if len(os.Getenv("SLEEP_TIME")) != 0 {
			sleepTime = os.Getenv("SLEEP_TIME")
		}

		sleepTimeDuration, _ := time.ParseDuration(sleepTime)
		time.Sleep(sleepTimeDuration)
	}
}
