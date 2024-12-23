package main

import (
	"log"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	log := log.New(os.Stderr, "", 0)
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Current time: %s", time.String())
}
