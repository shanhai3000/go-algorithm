package algorithm

import "algo/algorithm/util"

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}
func quickSort(arr []int, l int, r int) {
	if l < r {
		p := partition(arr, l, r)
		quickSort(arr, l, (*p)[0]-1)
		quickSort(arr, (*p)[1]+1, r)
	}
}

func partition(arr []int, l int, r int) *[2]int {
	pivotL := l - 1
	pivotR := r
	for ; l < pivotR; {
		if arr[l] < arr[r] {
			pivotL++
			util.Swap(arr, pivotL, l)
			l++
		} else if arr[l] > arr[r] {
			pivotR--
			util.Swap(arr, l, pivotR)
		} else {
			l++
		}
	}
	util.Swap(arr, pivotR, r)
	pivotL++
	return &[2]int{pivotL, pivotR}
}


