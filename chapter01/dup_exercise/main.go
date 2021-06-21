package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	occur_files := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, occur_files)
	} else {
		for _, arg := range files {
			// The function os.Open returns two values
			// The first is an open file (*os.File) that is used in subsequent reads by the Scanner.
			// The second result of os.Open is a value of the built-in error type
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, occur_files)
			f.Close()
		}
	}
	
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf("Occured in: %s \n\n", occur_files[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, occur_files map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		// fmt.Println(f.Name())
		// occur_files[input.Text()] += f.Name() + " "
		// append(occur_files[input.Text()], "asd")
		if stringInSlice(f.Name(), occur_files[input.Text()]) == false {
			occur_files[input.Text()] = append(occur_files[input.Text()], f.Name())
		}
		if input.Text() == "break" { break }
	}
}

// Check string in slice or not 
func stringInSlice(a string, slice []string) bool {
	for _, b := range slice {
		if b == a {
			return true
		}
	}
	return false
}