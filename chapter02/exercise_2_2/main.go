package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Measurement interface {
	String() string
}

type Distance struct {
	meters float64
}

func FromFeet(f float64) Distance {
	return Distance{f * 0.3048}
}

func FromMeters(m float64) Distance {
	return Distance{m}
}

func (d Distance) String() string {
	return fmt.Sprintf("%.3gm = %.3gft", d.meters, d.Feet())
}

func (d Distance) Meters() float64 {
	return d.meters
}

func (d Distance) Feet() float64 {
	return d.meters / 0.3048
}

type Temperature float64

func FromCelcius(c float64) Temperature {
	return Temperature(c)
}

func FromFarenheit(f float64) Temperature {
	return Temperature((f * 5 / 9) - 32)
}

func (t Temperature) String() string {
	return fmt.Sprintf("%.3gC = %.3gF", t.Celcius(), t.Farenheit())
}

func (t Temperature) Celcius() float64 {
	return float64(t)
}

func (t Temperature) Farenheit() float64 {
	return float64((t * 9 / 5) + 32)
}

func newMeasurement(f float64, unit string) (Measurement, error) {
	unit = strings.ToLower(unit)
	switch unit {
	case "m":
		return FromMeters(f), nil
	case "ft":
		return FromFeet(f), nil
	case "c":
		return FromCelcius(f), nil
	case "f":
		return FromFarenheit(f), nil
	default:
		fmt.Println("Default convert from meters.")
		return FromMeters(f), nil
	}
}

func printMeasurement(value string, unit string) {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatalf("%v isn't a number.", value)
	}

	newValue, err := newMeasurement(f, unit)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(newValue)
}

var unit = flag.String("u", "m", "unit")

func main() {
	flag.Parse()
	if len(os.Args) == 2 {
		printMeasurement(os.Args[1], *unit)
	} else if len(os.Args) >= 3 {
		for _, arg := range os.Args[2:] {
			printMeasurement(arg, *unit)
		}
	} 
}
