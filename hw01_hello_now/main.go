package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

const (
	layout  = "2006-01-02 15:04:05 +0000 UTC"
	ntpHost = "0.beevik-ntp.pool.ntp.org"
)

func main() {
	currentTime := time.Now().Format(layout)
	fmt.Printf("current time: %s\n", currentTime)

	ntpTime, err := ntp.Time(ntpHost)
	if err != nil {
		log.Fatalf("ntp time error: %s \n", err)
		os.Exit(1)
	}
	fmt.Printf("exact time: %s\n", ntpTime.Format(layout))
}
