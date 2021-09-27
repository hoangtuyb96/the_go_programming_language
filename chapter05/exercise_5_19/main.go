package main

import (
	"fmt"
)

func weird() (ret string) {
	defer func() {
		recover()
		ret = "recover done"
	}()
	panic("raise exception")
}

func main() {
	fmt.Println(weird())
}
