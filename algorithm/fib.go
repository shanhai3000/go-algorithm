package algorithm

/**
let's say fib sequence is started with 0, 1, 1, 2, 3, 5, 8, ... and so on
out : 44.692877843s (45.7265 in c++11)
*/
func Fib(n int) uint64{
	if n == 1{
		return 0
	}else if n == 2 || n == 3{
		return 1
	}

	return Fib(n-1) + Fib(n -2)
}