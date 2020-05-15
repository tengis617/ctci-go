package chapter1

/*
There are three types of edits that can be performed on strings:
insert a character, remove a character, or replace a character.
Given two strings, write a function to check if
they are one edit (or zero edits) away.
*/

func isOneEditAway(a, b string) bool {
	if len(a) == len(b) {
		return isOneReplaceAway(a, b)
	}
	if len(a)+1 == len(b) {
		return isOneEditInsertAway(a, b)
	}
	if len(a) == len(b)+1 {
		return isOneEditInsertAway(b, a)
	}
	return false
}

// isOneReplaceAway checks if a and b have only 1 char diff.
func isOneReplaceAway(a, b string) bool {
	var hasDifference bool
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			// already has 1 diff
			if hasDifference {
				return false
			}
			hasDifference = true
		}
	}
	return true
}

// isOneEditInsertAway checks if s1 is 1 insert away from s2.
// Note: len(s2) > len(s1)
func isOneEditInsertAway(s1, s2 string) bool {
	var index1, index2 int
	for {
		if index1 >= len(s1) || index2 >= len(s2) {
			return true
		}
		if s1[index1] != s2[index2] {
			if index1 != index2 {
				return false
			}
			index2++
		} else {
			index1++
			index2++
		}
	}
}
