package collection

import "algo/object"

type Stack struct {
	head  *Node
	next  *Node
	_size int
}

func (s *Stack) Push(o object.Object) {
	newNode := Node{o, nil}
	newNode.next = s.head
	s.head = &newNode
	s._size++
}
func (s *Stack) Pop() object.Object {
	if s.head != nil {
		t := s.head
		s.head = s.next
		s._size--
		return t.val
	}

	return nil
}

func (s *Stack) Empty() bool {
	return s._size == 0
}

func (s *Stack) Size() int {
	return s._size
}
