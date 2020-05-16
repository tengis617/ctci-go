package chapter1

import (
	"strconv"
	"strings"
)

/*
String Compression:
Implement a method to perform basic string compression
using the counts of repeated characters. For example, the string aabcccccaaa
would become a2b1c5a3. If the "compressed" string would not become smaller
than the original string, your method should return the original string.
You can assume the string has only uppercase and lowercase letters (a - z).
*/

// compress returns compressed string
// using counts of repeated characters.
// if compressed string is longer than original, it will
// return the original string.
func compress(s string) string {
	compLen := getCompressionLen(s)
	sLen := len(s)
	if compLen >= sLen {
		return s
	}

	var b strings.Builder
	b.Grow(compLen)
	var count int // keeps count of current char
	for i := range s {
		count++
		// if next char not the same
		// or its the last character of the string
		if i+1 >= sLen || s[i] != s[i+1] {
			b.WriteByte(s[i])
			b.WriteString(strconv.Itoa(count))
			count = 0
		}
	}
	return b.String()
}

// getCompressionLen computes the length of the compressed string
func getCompressionLen(s string) (length int) {
	var count int
	for i := range s {
		count++
		if i+1 >= len(s) || s[i] != s[i+1] {
			// we want to get the length of the count
			// ex: if count is 20 then we will add string "2" and "0"
			length += 1 + len(strconv.Itoa(count))
			count = 0
		}
	}
	return
}
