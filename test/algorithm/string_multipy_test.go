package algorithm

import (
	"algo/algorithm"
	"algo/algorithm/util"
	"testing"
	"time"
)

func TestStringMultiply(t *testing.T) {
	begin := time.Now()
	r := algorithm.StringMultiply("696231234325654756654321456786543", "3567489357483957483924738947380")
	use := time.Since(begin)
	util.DD(use)
	util.DD(r)
}
