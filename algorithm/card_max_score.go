package algorithm

/**
输入：cardPoints = [1,2,3,4,5,6,1], k = 3
输出：12
解释：第一次行动，不管拿哪张牌，你的点数总是 1 。但是，先拿最右边的卡牌将会最大化你的可获得点数。最优策略是拿右边的三张牌，最终点数为 1 + 6 + 5 = 12 。
 */
func CardMaxScore(cardPoints []int, k int) int {
	for i := 1; i < k && i < len(cardPoints); i++ {
		cardPoints[i] += cardPoints[i-1]
	}
	if k >= len(cardPoints) {
		return cardPoints[len(cardPoints)-1]
	}

	l := 0
	max := cardPoints[k-1]
	for i := len(cardPoints) - 1; i >= len(cardPoints)-k; i-- {
		l = k - (len(cardPoints) - i) - 1
		if l >= 0 {
			cardPoints[l+1] -= cardPoints[l]
			max = MaxInts(max, cardPoints[l]+cardPoints[i])
		} else {
			max = MaxInts(max, cardPoints[i])
		}
		cardPoints[i-1] += cardPoints[i]
	}

	return max
}
func MaxInts(i int, j int) int {
	if i >= j {
		return i
	}
	return j
}
