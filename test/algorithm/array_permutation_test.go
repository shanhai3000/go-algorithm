package algorithm

import (
	"algo/algorithm"
	"algo/algorithm/util"
	"testing"
	"time"
)

func TestFullPermutation(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	begin := time.Now()
	res := algorithm.FullPermutation(nums)
	util.DD(time.Since(begin))
	util.DD(res)
}
