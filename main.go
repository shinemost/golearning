package main

import (
	"log"
	"time"
)

func main() {
	TickerDemo()
}

func TickerDemo() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Ticker tick.")
	}

}
