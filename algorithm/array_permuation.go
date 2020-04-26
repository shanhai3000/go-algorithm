package algorithm

func FullPermutation(nums []int) [][]int {
	var ans [][]int
	var currentPermutation []int
	var visited = make([]bool, len(nums))
	depth := 0
	dfs(nums, &ans, visited, currentPermutation, depth)
	return ans
}
func dfs(nums []int, ans *[][]int, visited []bool, currentPermutation []int, depth int) {
	if depth == len(nums) {
		tmp := make([]int, len(currentPermutation))
		copy(tmp, currentPermutation)
		*ans = append(*ans, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !visited[i] {
			visited[i] = true
			depth++
			currentPermutation = append(currentPermutation, nums[i])
			dfs(nums, ans, visited, currentPermutation, depth)
			depth--
			visited[i] = false
			currentPermutation = currentPermutation[:len(currentPermutation)-1]
		}
	}
}
