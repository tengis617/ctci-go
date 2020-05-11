package ctcigo

// Question1: What is the runtime?
func product(a, b int) int {
	sum := 0
	for i := 0; i < b; i++ { // O(b) operation
		sum += a // constant O(1) operation
	}
	return sum
}

// Answer: O(b)

// Question2:
func power(a, b int) int {
	if b < 0 {
		return 0 // error
	} else if b == 0 {
		return 1
	} else {
		return a * power(a, b-1) //
	}
}
// Answer: O(b)
// a = 1, b = 1 --> power(1, 1), power(1, 0) times: 2
// a = 2, b = 1 --> power(2, 1), power(2, 0) times: 2
// looks like a is not increasing work 
// a = 1, b = 2 --> power(1, 2), power(1, 1), power(1, 0)  times: 3
// a, b --> power(a, b), power(a, b - 1), power(a, b - 2) ... power(a, b - b)  times: b