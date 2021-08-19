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

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	for i := 0; i < n; i++ {
		if (n-i)%3 == 0 && i != 0 {
			buf.WriteString(",")
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

//!-
