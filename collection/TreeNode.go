package collection

import "algo/object"

type TreeNode struct {
	val object.Object
	left *TreeNode
	right *TreeNode
}

type Node struct {
	val object.Object
	next *Node
}