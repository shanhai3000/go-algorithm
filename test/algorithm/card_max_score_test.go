package algorithm

import (
	"algo/algorithm"
	"algo/algorithm/util"
	"testing"
)

func TestCardMaxScore(t *testing.T) {
	arr := []int{2,2,2}
	util.DD(algorithm.CardMaxScore(arr, 2))
}
