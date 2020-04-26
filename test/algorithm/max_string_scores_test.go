package algorithm

import (
	"algo/algorithm"
	"algo/algorithm/util"
	"testing"
)

func TestMaxStringScores(t *testing.T){
	s := "00111"
	util.DD(algorithm.MaxStringScores(s))
}
