package collection

type Node struct {
	value interface{}
	next  *Node
}

type Stack struct {
	head  *Node
	_size int
}

func (s *Stack) Push(value interface{}) {
	newNode := Node{value, nil}
	newNode.next = s.head
	s.head = &newNode
	s._size++
}
func (s *Stack) Pop() interface{} {
	if s.head != nil {
		t := s.head
		s.head = s.head.next
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
