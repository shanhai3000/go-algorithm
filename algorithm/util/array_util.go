package util


func Swap(arr []int, i int, j int) {
	if arr[i] != arr[j] {
		arr[i] ^= arr[j]
		arr[j] ^= arr[i]
		arr[i] ^= arr[j]
	}
}
