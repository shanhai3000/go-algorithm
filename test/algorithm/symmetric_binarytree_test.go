package algorithm

import (
	"algo/algorithm"
	"algo/collection/node"
	"testing"
	"time"
)

/**
 [1,2,2,3,4,4,3] symmetric。
    1
   / \
  2   2
 / \ / \
3  4 4  3
*/

func TestIsSymmetric1(t *testing.T) {
	root := new(node.TreeNode)
	root.Val = 1
	root.Left = &node.TreeNode{Val: 2}
	root.Right = &node.TreeNode{Val: 2}
	root.Left.Left = &node.TreeNode{Val: 3}
	root.Left.Right = &node.TreeNode{Val: 4}
	root.Right.Left = &node.TreeNode{Val: 4}
	root.Right.Right = nil

	begin := time.Now()
	ret := algorithm.IsSymmetric0(root)
	if ret {
		t.Log("IsSymmetric0() Ok, use ", time.Since(begin))
	} else {
		t.Error("Error")
	}

	begin = time.Now()
	ret = algorithm.IsSymmetric1(root)
	if ret {
		t.Log("IsSymmetric1() Ok, use ", time.Since(begin))
	} else {
		t.Error("Error")
	}
}
