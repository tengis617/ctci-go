package chapter1

/*
Rotate Matrix: Given an image represented by an NxN matrix,
where each pixel in the image is 4 bytes,
write a method to rotate the image by 90 degrees. (an you do this in place?
*/

// rotateMatrix rotates NxN matrix by 90 degrees.
// Runtime O(N^2)
// Memory O(N)
func rotateMatrix(m [][]int) [][]int {
	for layer := 0; layer < len(m)/2; layer++ {
		first := layer
		last := len(m) - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			top := m[first][i]
			// left -> top
			m[first][i] = m[last-offset][first]
			// bottom -> left
			m[last-offset][first] = m[last][last-offset]
			// right -> bottom
			m[last][last-offset] = m[i][last]
			// top -> right
			m[i][last] = top
		}
	}

	return m
}
