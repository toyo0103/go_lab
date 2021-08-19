package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("  %t\n", isReverse(os.Args[1], os.Args[2]))
}

func isReverse(s1 string, s2 string) bool {
	n := len(s1)
	if n != len(s2) {
		return false
	}

	for i := 0; i < n; i++ {
		if s1[i] != s2[n-(i+1)] {
			return false
		}
	}

	return true
}

//!-
