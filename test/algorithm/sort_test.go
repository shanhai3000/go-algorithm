package algorithm

import (
	"algo/algorithm"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {
	var arr []int
	for i := 0; i < 100000000; i++ {
		arr = append(arr, rand.Intn(1000))
	}
	begin := time.Now()
	algorithm.QuickSort(arr)
	use := time.Since(begin)
	fmt.Println("use ", use)
}/* Sort 100000000 elements using 5s */

func TestHeapSort(t *testing.T) {
	var arr []int
	for i := 0; i < 100000000; i++ {
		arr = append(arr, rand.Intn(1000))
	}
	begin := time.Now()
	algorithm.HeapSort(arr)
	use := time.Since(begin)
	fmt.Println("use ", use)
}/* Sort 100000000 elements using 31.8s */

