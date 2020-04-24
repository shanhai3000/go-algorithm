package algorithm

import (
	"algo/collection/node"
	"algo/collection/stack"
	"fmt"
)

/**
 [1,2,2,3,4,4,3] symmetric。
    1
   / \
  2   2
 / \ / \
3  4 4  3
 [1,2,2,null,3,null,3] not symmetric:
    1
   / \
  2   2
   \   \
   3    3
*/

func IsSymmetric0(root *node.TreeNode) bool {
	return isSymmetric0(root, root)
}
func isSymmetric0(l *node.TreeNode, r *node.TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}

	return l.Val == r.Val && isSymmetric0(l.Left, r.Right) && isSymmetric0(l.Right, r.Left)
} /* recursion version*/

func IsSymmetric1(root *node.TreeNode) bool {
	if root == nil {
		return true
	}
	flag := true
	s := new(stack.Stack)
	s.Push(root.Left)
	s.Push(root.Right)
	for ; !s.Empty(); {
		l0 := s.Pop().(*node.TreeNode)
		l1 := s.Pop().(*node.TreeNode)
		if l0 != nil && l1 != nil {
			if l0.Val != l1.Val {
				fmt.Println("l0=", l0.Val, "l1=", l1.Val)
				flag = false
				break
			}
			s.Push(l0.Left)
			s.Push(l1.Right)
			s.Push(l0.Right)
			s.Push(l1.Left)
		} else if l0 != nil || l1 != nil {
			flag = false
			break
		}
	}

	return flag
}
