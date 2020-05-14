package chapter1

// Implement an algorithm to determine if a string
// has all unique characters. What if you cannot use
// any additional Data Structures?

// time complexity: O(N) => where N is len of string
// space complexity: O(1)
func isUniqueHashTable(s string) bool {
	// if we assume we are using ASCII characters only:
	if len(s) > 128 {
		return false
	}
	exists := make(map[rune]bool, len(s))
	for _, c := range s {
		if exists[c] {
			return false
		}
		exists[c] = true
	}
	return true
}

// time complexity: O(N^2)
// space complexity: O(1)
func isUniqueLoop(s string) bool {
	runes := []rune(s)
	// fmt.Printf("string: %s \n", s)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if i == j {
				continue
			}
			if runes[i] == runes[j] {
				// fmt.Printf("%d: %c , %d: %c \n", i, runes[i], j, runes[j])
				return false
			}
		}
	}
	return true
}

// TODO: implement bit vector solution
