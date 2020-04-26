package algorithm

import "algo/algorithm/util"

/**
	"000111" ---->  "000" + "111"
		score = 3 + 3 = 6
 */
func MaxStringScores(s string)int{
	ans, c0, c1 := 0,0,0
	for i:=0; i < len(s);i++{
		if s[i] == '1'{
			c1++
		}
	}
	for i := 0;i < len(s)-1; i++{
		if s[i] == '0'{
			c0++
		}else{
			c1--
		}
		ans = util.MaxInts(ans, c0+c1)
	}
	return ans
}
