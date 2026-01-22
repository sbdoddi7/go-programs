package main

import "fmt"

// Reverse a string in Go in place O(1) Space and return the reversed string.

func reverseString(s []byte) string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func main() {
	fmt.Println(reverseString([]byte("Hello, World!")))
}
