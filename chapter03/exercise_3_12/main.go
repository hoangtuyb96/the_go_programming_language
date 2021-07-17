package main

import (
	"fmt"
	"os"
)

func main() {
	a := os.Args[1]
	b := os.Args[2]

	if isAnagram(a, b) {
		fmt.Println("They are anagrams")
	} else {
		fmt.Println(("They aren't anagrams"))
	}
}

func isAnagram(a, b string) bool {
	aFreq := make(map[rune]int)
	for _, c := range a {
		aFreq[c]++
	}
	bFreq := make(map[rune]int)
	for _, c := range b {
		bFreq[c]++
	}
	for k, v := range aFreq {
		if bFreq[k] != v {
			return false
		}
	}
	for k, v := range bFreq {
		if aFreq[k] != v {
			return false
		}
	}
	return true
}
