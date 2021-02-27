//+build release

package main

import (
	"log"
	"os"
	"time"
)

func init () {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	SetMonitorTick(15 * time.Minute)
}
