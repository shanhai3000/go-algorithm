package util

import (
	"fmt"
	"os"
)

func DD(a ...interface{}) {
	fmt.Fprintln(os.Stdout, a...)
}

func MaxInts(i int, j int) int {
	if i >= j {
		return i
	}
	return j
}
