package main

import "fmt"
import "github.com/the_go_programming_language/chapter02/exercise_2_1/tempconv"

func main() {
	freezingK := tempconv.CToK(tempconv.FreezingC)
	fmt.Printf("Freezing in %g°K\n", freezingK)

	tempInK := 255.00
	fmt.Printf("%g°K is %g°C\n", tempInK, tempconv.KToC(tempconv.Kelvin(tempInK)))
}
