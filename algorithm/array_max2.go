package algorithm

func ArrayMax2(arr []int) [2]int {
	max, med := arr[0], arr[1]
	if med > max {
		med ^= max
		max ^= med
		med ^= max
	}
	lo, hi := 2, len(arr)
	for ; lo < hi; lo++ {
		if arr[lo] > max {
			med = max
			max = arr[lo]
		}
		if arr[lo] > med && arr[lo] < max {
			max = arr[lo]
		}
	}
	return [2]int{max, med}
}
