package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Print("Index:" + strconv.Itoa(i) + " / ")
		fmt.Println("Value: " + os.Args[i])
	}
}