package algorithm

/**
 n 名士兵站成一排。每个士兵都有一个 独一无二 的评分 rating 。

每 3 个士兵可以组成一个作战单位，分组规则如下：

从队伍中选出下标分别为 i、j、k 的 3 名士兵，他们的评分分别为 rating[i]、rating[j]、rating[k]
作战单位需满足： rating[i] < rating[j] < rating[k] 或者 rating[i] > rating[j] > rating[k] ，其中  0 <= i < j < k < n
请你返回按上述条件可以组建的作战单位数量。每个士兵都可以是多个作战单位的一部分。
示例 1：

输入：rating = [2,5,3,4,1]
输出：3
解释：我们可以组建三个作战单位 (2,3,4)、(5,4,1)、(5,3,1) 。
示例 2：

输入：rating = [2,1,3]
输出：0
解释：根据题目条件，我们无法组建作战单位。
示例 3：

输入：rating = [1,2,3,4]
输出：4
*/
func NumTeams(rating []int) int {
	sum := 0
	for i := 0; i < len(rating); i++ {
		numTeams0(rating, i, 0, 1, &sum)
	}
	return sum
}
func numTeams0(rating []int, i int, delt int, curNum int, sum *int) {
	if curNum == 3 {
		*sum++
		return
	}
	for k := i + 1; k < len(rating); k++ {
		if delt == 0 {
			if rating[k] > rating[i] {
				numTeams0(rating, k, delt+1, curNum+1, sum)
			} else if rating[k] < rating[i] {
				numTeams0(rating, k, delt-1, curNum+1, sum)
			}
		} else {
			if delt > 0 && rating[k] > rating[i] {
				numTeams0(rating, k, delt, curNum+1, sum)
			} else if delt < 0 && rating[k] < rating[i] {
				numTeams0(rating, k, delt, curNum+1, sum)
			}
		}
	}
}
