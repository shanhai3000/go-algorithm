package algorithm

import (
	"algo/collection/node"
	"fmt"
)

func PreOrderTraverse0(root *node.TreeNode){
	// l-l-l-l-l-l-root-r
	if root != nil {
		fmt.Print(root.Val)
		PreOrderTraverse0(root.Left)
		PreOrderTraverse0(root.Right)
	}
}/*recursion version*/