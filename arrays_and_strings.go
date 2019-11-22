package main

import (
	"strings"
)

/*
Is Unique:

Implement an algorithm to determine if a string has all unique characters.
What if you cannot use additional data structures?
*/

// isUniqueHashMap checks whether a string has all unique characters using hash tables.
func isUniqueHashMap(s string) bool {
	charSet := make(map[rune]bool, len(s))
	for _, c := range s {
		if charSet[c] {
			return false
		}
		charSet[c] = true
	}
	return true
}

/*
Check Permutation:

Given two strings, write a method to decide if
one is a permutation of the other.
*/
// isPermutation checks if a string is a permutation of another.
// runs in O(A + B)
// space O(1) if we do ASCII
func isPermutation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	charSet := make(map[rune]int)

	for _, c := range a {
		charSet[c] += 1
	}

	for _, c := range b {
		charSet[c] -= 1
		if charSet[c] < 0 {
			return false
		}
	}
	return true
}

/*
URLify:

Write a method to replace all spaces in a string with '%20:
You may assume that the string has sufficient space at the end
to hold the additional characters, and that you are given the
"true" length of the string.
*/
func URLify(s string, l int) string {
	var sb strings.Builder
	sb.Grow(l)
	for i, c := range s {
		if i >= l {
			break
		}
		if c == ' ' {
			sb.WriteString("%20")
		} else {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}
