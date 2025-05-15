// 0036. Valid Sudoku
// https://neetcode.io/problems/valid-sudoku

package medium

// 1. Brute Force
// Time Complexity: O(n^2)
// Space Complexity: O(n^2)
func ValidSudokuSolution1(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		seen := make(map[byte]bool)
		for i := 0; i < 9; i++ {
			if board[row][i] == '.' {
				continue
			}
			if seen[board[row][i]] {
				return false
			}
			seen[board[row][i]] = true
		}
	}

	for col := 0; col < 9; col++ {
		seen := make(map[byte]bool)
		for i := 0; i < 9; i++ {
			if board[i][col] == '.' {
				continue
			}
			if seen[board[i][col]] {
				return false
			}
			seen[board[i][col]] = true
		}
	}

	for box := 0; box < 9; box++ {
		seen := make(map[byte]bool)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				row := (box/3)*3 + i
				col := (box%3)*3 + j
				if board[row][col] == '.' {
					continue
				}
				if seen[board[row][col]] {
					return false
				}
				seen[board[row][col]] = true
			}
		}
	}

	return true
}

// 2. Hash Set (One Pass)
// Time Complexity: O(n^2)
// Space Complexity: O(n^2)
func ValidSudokuSolution2(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			cell := board[row][col]
			if cell == '.' {
				continue
			}

			box := (row/3)*3 + col/3

			if rows[row][cell] || cols[col][cell] || boxes[box][cell] {
				return false
			}

			rows[row][cell] = true
			cols[col][cell] = true
			boxes[box][cell] = true
		}
	}

	return true
}

// 3. Bitmask
// Time Complexity: O(n^2)
// Space Complexity: O(n)
func ValidSudokuSolution3(board [][]byte) bool {
	rows := make([]int, 9)
	cols := make([]int, 9)
	boxes := make([]int, 9)

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			cell := board[row][col]
			if cell == '.' {
				continue
			}

			cell -= '1'
			bit := 1 << cell
			box := (row/3)*3 + col/3

			if rows[row]&bit != 0 || cols[col]&bit != 0 || boxes[box]&bit != 0 {
				return false
			}

			rows[row] |= bit
			cols[col] |= bit
			boxes[box] |= bit
		}
	}

	return true
}
