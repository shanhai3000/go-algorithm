package algorithm

/**
board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false
*/
func WordExist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	if len(word) == 0 {
		return true
	}
	exists := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		t := make([]bool, len(board[0]))
		exists[i] = t
	}
	if find(board, word, 0, 0, 0, exists) {
		return true
	}
	return false
}
func find(board [][]byte, word string, curIdx int, row int, col int, exists [][]bool) bool {
	if curIdx == len(word) {
		return true
	}
	curChar := word[curIdx]
	if curIdx == 0 {
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				if board[i][j] == curChar {
					exists[i][j] = true
					if find(board, word, curIdx+1, i, j, exists) {
						return true
					}
					exists[i][j] = false
				}
			}
		}
		return false
	} else {
		if row-1 >= 0 {
			if board[row-1][col] == curChar && !exists[row-1][col] {
				exists[row-1][col] = true
				if find(board, word, curIdx+1, row-1, col, exists) {
					return true
				}
				exists[row-1][col] = false
			}
		}
		if row+1 < len(board) {
			if board[row+1][col] == curChar && !exists[row+1][col] {
				exists[row+1][col] = true
				if find(board, word, curIdx+1, row+1, col, exists) {
					return true
				}
				exists[row-1][col] = false
			}
		}
		if col-1 >= 0 {
			if board[row][col-1] == curChar && !exists[row][col-1] {
				exists[row][col-1] = true
				if find(board, word, curIdx+1, row, col-1, exists) {
					return true
				}
				exists[row][col-1] = false
			}
		}
		if col+1 < len(board[0]) {
			if board[row][col+1] == curChar && !exists[row][col+1] {
				exists[row][col+1] = true
				if find(board, word, curIdx+1, row, col+1, exists) {
					return true
				}
				exists[row][col+1] = false
			}
		}
		return false
	}

}
