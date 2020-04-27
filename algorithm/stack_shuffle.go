package algorithm

import "algo/collection/stack"

/**
栈混洗
给定一个序列 [1,2,3]，判断 另一个序列是否为该序列的一个栈混洗
*/
func StackShuffle(nums []int, seq []int) {
	src, dst := new(stack.Stack), new(stack.Stack)
	for i := len(nums) - 1; i >= 0; i-- {
		src.Push(nums[i])
	}
	c, seqIdx, subSeq := 0, 0, true
	for i := 0; i < len(nums); i++ {
		if !subSeq {
			break
		}
		if dst.Empty() {
			dst.Push(src.Pop())
		}
		for !dst.Empty() {
			c = dst.Peek().(int)
			if c == seq[seqIdx] { //栈顶元素与期望的序列元素相等，出栈
				dst.Pop()
				seqIdx++
				break //为避免栈已空 出现异常指针访问
			} else {
				if src.Empty() {
					subSeq = false
					break
				}
				dst.Push(src.Pop())
			}
		}
	}
}
