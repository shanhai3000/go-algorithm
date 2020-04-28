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

func PreOrderTraverse1(root *node.TreeNode, visitor func(*node.TreeNode)) {
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
/**
		10
     5      15
   1   6   12  16
*/
func MidOrderTraverse0(root *node.TreeNode, visitor func(*node.TreeNode)){
	if root != nil 	{
		if root.Left != nil {
			MidOrderTraverse0(root.Left, visitor)
		}
		visitor(root)
		if root.Right != nil {
			MidOrderTraverse0(root.Right, visitor)
		}
	}
}/*recursion version*/

func MidOrderTraverse1(root *node.TreeNode, visitor func(*node.TreeNode)){
	s := new(stack.Stack)
	x := root
	for  {
		goAlongLeftBranch(x, s)
		if s.Empty(){
			break
		}
		x = s.Pop().(*node.TreeNode)
		visitor(x)
		x = x.Right//可能为空，留意处理手法，以及这样写的目的
	}
}/*	iteration version*/
 func goAlongLeftBranch(root *node.TreeNode, s *stack.Stack){
 	for root!= nil{
 		s.Push(root)
 		root=root.Left
	}
 }