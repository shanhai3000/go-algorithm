package algorithm

import (
	"algo/collection/node"
	"algo/collection/stack"
	"fmt"
)
func PrintNode(node *node.TreeNode){
	fmt.Println(node.Val)
}

func PreOrderTraverse0(root *node.TreeNode, visitor func(*node.TreeNode)){
	// l-l-l-l-l-l-root-r
	if root != nil {
		visitor(root)
		PreOrderTraverse0(root.Left, visitor)
		PreOrderTraverse0(root.Right, visitor)
	}
}/*recursion version*/

func PreOrderTraverse1(root *node.TreeNode, visitor func(treeNode *node.TreeNode)) {
	s := new(stack.Stack)
	s.Push(root)
	for ; !s.Empty(); {
		t := s.Pop().(*node.TreeNode)
		visitor(t)
		if t.Right != nil{
			s.Push(t.Right)
		}
		if t.Left != nil {
			s.Push(t.Left)
		}
	}
}
