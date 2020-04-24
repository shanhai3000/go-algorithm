package algorithm

import (
	"algo/collection/node"
	"fmt"
	"testing"
	"time"
)

/**
    6
   / \
  4   8
 / \ / \
3  5 7  9
*/
func TestPreOrderTraverse0(t *testing.T) {
	root := new(node.TreeNode)
	root.Val = 6
	root.Left = &node.TreeNode{Val: 4}
	root.Right = &node.TreeNode{Val: 8}
	root.Left.Left = &node.TreeNode{Val: 3}
	root.Left.Right = &node.TreeNode{Val: 5}
	root.Right.Left = &node.TreeNode{Val: 7}
	root.Right.Right = &node.TreeNode{Val: 9}
	begin := time.Now()
	PreOrderTraverse0(root, PrintNode)
	fmt.Println(time.Since(begin))
	begin = time.Now()
	PreOrderTraverse1(root, PrintNode)
	fmt.Println(time.Since(begin))
}
