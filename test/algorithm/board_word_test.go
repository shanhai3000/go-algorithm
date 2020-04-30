package algorithm

import (
	"algo/algorithm"
	"algo/algorithm/util"
	"testing"
)

func TestBoardWord(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABS"
	res := algorithm.WordExist(board, word)
	util.DD(res)
}
