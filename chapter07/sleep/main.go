package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

// go run main.go -period 50ms
// go run main.go -period 2m30s
// go run main.go -period 1.5h
// go run main.go -period "1 day" // invalid value "1 day" for flag -period: time: invalid duration 1 day
