package chapter1

/*
Zero Matrix:
Write an algorithm such that if an element in an MxN matrix is 0,
its entire row and column are set to 0.
*/

func zeroMatrix(m [][]int) [][]int {
	var rowHasZero, colHasZero bool

	// check first row
	for j := 0; j < len(m[0]); j++ {
		if m[0][j] == 0 {
			rowHasZero = true
			break
		}
	}

	// check first col
	for i := 0; i < len(m); i++ {
		if m[i][0] == 0 {
			colHasZero = true
			break
		}
	}

	// check for zeros
	for i := 1; i < len(m); i++ {
		for j := 1; j < len(m[0]); j++ {
			if m[i][j] == 0 {
				m[i][0] = 0
				m[0][j] = 0
			}
		}
	}

	// clear remaining row based on first col
	for i := 1; i < len(m); i++ {
		if m[i][0] == 0 {
			clearRow(m, i)
		}
	}

	// clear remaining col based on first col
	for j := 1; j < len(m[0]); j++ {
		if m[0][j] == 0 {
			clearCol(m, j)
		}
	}

	// clear first row
	if rowHasZero {
		clearRow(m, 0)
	}
	// clear first col
	if colHasZero {
		clearCol(m, 0)
	}

	return m
}

func clearRow(m [][]int, row int) {
	for j := 0; j < len(m[row]); j++ {
		m[row][j] = 0
	}
}

func clearCol(m [][]int, col int) {
	for i := 0; i < len(m); i++ {
		m[i][col] = 0
	}
}
