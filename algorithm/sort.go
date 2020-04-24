package algorithm

import (
	"algo/algorithm/util"
)

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
} /* Sort 100000 elements using 16.057379ms */
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

/**
		1
      /    \
	 2       33
    / \     / \
   9   9   7   3
  / \ / \ / \ /  \
 5  5 8 3 3 2 5  16
一个n层的满二叉树，节点总数为2**n-1
imagine the array as a binary tree
自倒数第二层开始向上heapify
heapify0：
	-	和自己的字节点比较大小，并和较大者交换的过程
	-	交换完成后，子节点分别继续heapify
*/
func HeapSort(arr []int) {
	length := len(arr)
	for i := (length - 1) >> 1; i >= 0; i-- {
		percolateDown(arr, i, length)
	}
	//the max number is the root node
	for i := length - 1; i > 0; i-- {
		util.Swap(arr, i, 0)
		percolateDown(arr, 0, i)
	}
}

func heapify0(arr []int, i int, len int) {
	var newI int
	if i < len {
		newI = i<<1 + 1
		if newI < len && arr[newI] > arr[i] {
			util.Swap(arr, newI, i)
			heapify0(arr, newI, len)
		}
		newI = i<<1 + 2
		if newI < len && arr[newI] > arr[i] {
			util.Swap(arr, newI, i)
			heapify0(arr, newI, len)
		}
	}
} /*recursion version. Sort 100000 elements, using  260ms */

func heapify1(arr []int, i int, len int) {
	var newI int
	for ; i < len; i++ {
		newI = i<<1 + 1
		if newI < len && arr[newI] > arr[i] {
			util.Swap(arr, newI, i)
		}
		newI = i<<1 + 2
		if newI < len && arr[newI] > arr[i] {
			util.Swap(arr, newI, i)
		}
	}
}

/**
下滤
*/
func percolateDown(arr []int, i int, len int) {
	var childIdx int
	var nextIdx int
	var move bool
	for ; i < len; {
		move = false
		childIdx = i<<1 + 1
		if childIdx < len && arr[childIdx] > arr[i] {
			nextIdx = childIdx
			move = true
		}
		childIdx = i<<1 + 2
		if childIdx < len && arr[childIdx] > arr[i] && arr[childIdx] > arr[childIdx-1] {
			nextIdx = childIdx
			move = true
		}
		if !move {
			break
		}
		util.Swap(arr, i, nextIdx)
		i = nextIdx
	}
}
