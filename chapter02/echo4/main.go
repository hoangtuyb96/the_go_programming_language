package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println("asdasd")
	}
}

// go run main.go a bc def
// go run main.go -s / a bc def
// go run main.go -n a bc def
// go run main.go -help
