package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer

	pre := len(s) % 3
	if pre == 0 {
		pre = 3
	}

	buf.WriteString(s[:pre])
	for i := pre; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i:i+3])
	}
	return buf.String()
}
