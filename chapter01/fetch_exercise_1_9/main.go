package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// check prefix of url
		if strings.HasPrefix(url, "http://") != true {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v \n", err)
			os.Exit(1)
		}

		fmt.Printf("Response status: %s \n", resp.Status)
		fmt.Printf("Response status code: %d \n", resp.StatusCode)
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v \n", err)
			os.Exit(1)
		}
	}
}

// go run .\main.go gopl.io
