package chapter1

/*
URLify:
Write a method to replace all spaces in a string with '%20'.
You may assume that the string has sufficient space at the end to hold the additional characters,
and that you are given the "true" length of the string.
*/

// urlify replaces all spaces in the given string with %20.
// trueLength is the true length of the string
// Time Complexity: O(N)
// Space Complexity: O(N)
// where N is num of chars in s
func urlify(s string, trueLength int) string {
	chars := []rune(s)
	// i will be used to point to the index to write at
	i := len(chars)
	// j is for keeping track of characters (starting from true length - 1) in reverse order
	for j := trueLength - 1; j >= 0; j-- {
		if chars[j] == ' ' {
			chars[i-1] = '0'
			chars[i-2] = '2'
			chars[i-3] = '%'
			i -= 3
		} else {
			chars[i-1] = chars[j]
			i--
		}
	}
	return string(chars)
}
