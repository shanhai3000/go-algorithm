package test

import (
	"algo/algorithm"
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{2, 1, 1, 5, 6, 3, 5}
	algorithm.QuickSort(arr)
	fmt.Println(arr)
}
