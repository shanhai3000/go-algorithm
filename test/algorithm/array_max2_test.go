package algorithm

import (
	"algo/algorithm"
	"algo/algorithm/util"
	"math/rand"
	"testing"
	"time"
)

func TestArrayMax2(t *testing.T) {
	var arr []int
	for i := 0; i < 1000000; i ++{
		arr = append(arr, rand.Intn(1000000))
	}
	arr = append(arr, 1000001)
	arr = append(arr, 1000011)
	begin := time.Now()
	ret := algorithm.ArrayMax2(arr)
	util.DD(time.Since(begin))
	util.DD(ret)
}
