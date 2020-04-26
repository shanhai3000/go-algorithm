package util

import "fmt"

func DD(in interface{}) {
	fmt.Println(in)
}

func MaxInts(i int, j int) int {
	if i >= j {
		return i
	}
	return j
}
