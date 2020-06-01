package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

const (
	ntpHost = "0.beevik-ntp.pool.ntp.org"
)

func main() {
	currentTime := time.Now().Round(0)
	fmt.Printf("current time: %s\n", currentTime.String())

	ntpTime, err := ntp.Time(ntpHost)
	if err != nil {
		log.Fatalf("ntp time error: %s \n", err)
	}
	exactTime := ntpTime
	fmt.Printf("exact time: %s\n", exactTime)
}
