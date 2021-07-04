package main

import "fmt"
import "github.com/the_go_programming_language/chapter02/tempconv0/tempconv"

func main() {
	fmt.Printf("%g\n", tempconv.BoilingC - tempconv.FreezingC) // "100" °C
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", boilingF - tempconv.CToF(tempconv.FreezingC)) // "180" °F

	c := tempconv.FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"; does not call String
	fmt.Println(float64(c)) // "100"; does not call String
}
