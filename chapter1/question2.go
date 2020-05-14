package chapter1

/*
Check Permutation
Given two strings, write a method to decide
if one is a permutation of the other.
*/

// isPermutation checks if one string is a permutation of the other.
// Time complexity = O(N) => where N is len of character
// Space complexity = O(1) => assuming there is a fixed number of characters (ASCII=128 etc.)
func isPermutation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	count := make(map[rune]int, 128)

	for _, c := range a {
		count[c] += 1
	}
	for _, c := range b {
		if count[c] < 1 {
			return false
		}
		count[c] -= 1
	}
	return true
}
