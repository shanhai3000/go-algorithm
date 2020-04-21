package collection

type Stack struct {
	head  *Node
	next  *Node
	_size int
}

func (s *Stack) Push(val interface{}) {
	newNode := Node{val, nil}
	newNode.next = s.head
	s.head = &newNode
	s._size++
}
func (s *Stack) Pop() interface{} {
	if s.head != nil {
		t := s.head
		s.head = s.next
		s._size--
		return t.value
	}

	return nil
}

func (s *Stack) Empty() bool {
	return s._size == 0
}

func (s *Stack) Size() int {
	return s._size
}
