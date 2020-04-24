package algorithm

import (
	"algo/algorithm"
	"fmt"
	"testing"
	"time"
)

func TestFib1(t *testing.T) {
	n := 200
	begin := time.Now()
	fmt.Println(algorithm.Fib1(n))
	fmt.Println("Fib1 use ", time.Since(begin))//out 4.64µs
	begin = time.Now()
	fmt.Println(algorithm.Fib2(n))
	fmt.Println("Fib2 use", time.Since(begin))//out  1.738µs
}
