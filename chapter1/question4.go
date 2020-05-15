package chapter1

import (
	"fmt"
	"strings"
)

/*
1.4 Palindrome Permutation
*/
func isPalindromePermutation(s string) bool {
	// cleanup
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")

	fmt.Println(s)
	charCount := make(map[rune]int)
	for _, c := range s {
		charCount[c] += 1
	}
	oddCount := 0
	for _, i := range charCount {
		if i%2 != 0 {
			oddCount += 1
		}
	}
	return oddCount <= 1
}
