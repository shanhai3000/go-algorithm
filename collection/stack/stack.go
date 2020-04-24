package stack

type (
	node struct {
		Value interface{}
		Next  *node
	}

	Stack struct {
		root  *node
		_size int
	}
)

func (s *Stack) Push(value interface{}) {
	newNode := node{value, nil}
	newNode.Next = s.root
	s.root = &newNode
	s._size++
}
func (s *Stack) Pop() interface{} {
	if s.root != nil {
		t := s.root
		s.root = s.root.Next
		s._size--
		return t.Value
	}

	return nil
}

func (s *Stack) Empty() bool {
	return s._size == 0
}

func (s *Stack) Size() int {
	return s._size
}
