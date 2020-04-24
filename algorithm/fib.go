package algorithm

/**
let's say fib sequence is started with 0, 1, 1, 2, 3, 5, 8, ... and so on
out : 44.692877843s (45.7265 in c++11)
*/
func Fib(n int) uint64{
	if n <= 1{
		return 0
	}else if n == 2 || n == 3{
		return 1
	}

	return Fib(n-1) + Fib(n -2)
}/* recursion version, bad */

func Fib1(n int) uint64	{
	arr := []uint64{0,1}
	if n < 2{
		return arr[n-1]
	}
	for i := 2; i < n; i++ {
		arr[1] = arr[0] + arr[1]
		arr[0] = arr[1] - arr[0]//two times operation, bad
	}
	return arr[1]
}

func Fib2(n int) uint64	{
	arr := []uint64{0,1}
	if n < 2{
		return arr[n-1]
	}
	for i := 2; i < n; i++ {
		t := arr[1]
		arr[1] = arr[0] + arr[1]
		arr[0] = t
	}
	return arr[1]
}/*best version*/