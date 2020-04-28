package algorithm

import (
	"algo/algorithm/util"
	"algo/collection/sort"
	"algo/collection/stack"
	"testing"
)

type MyInt int

func (i MyInt) CompareTo(t sort.Comparable) int {
	if i > t.(MyInt) {
		return 1
	} else if i == t.(MyInt) {
		return 0
	} else {
		return -1
	}
}
func TestSortedStack(t *testing.T) {
	s := new(stack.SortedStack)
	util.DD(s.Peek())
	s.Pop()
	util.DD(s.Peek())
	s.Push(MyInt(40))
	util.DD(s.Peek())
	s.Pop()
	util.DD(s.Peek())
	s.Push(MyInt(19))
	util.DD(s.Peek())
	s.Push(MyInt(44))
	util.DD(s.Peek())
	s.Pop()
	util.DD(s.Peek())
	s.Pop()
	util.DD(s.Peek())
	s.Push(MyInt(42))
	util.DD(s.Peek())
	s.Push(MyInt(8))
	util.DD(s.Peek())
	s.Push(MyInt(29))
	util.DD(s.Peek())
	s.Push(MyInt(25))
	util.DD(s.Peek())
	s.Pop()
	util.DD(s.Peek())
	s.Pop()
	util.DD(s.Peek())
	s.Push(MyInt(52))
	util.DD(s.Peek())
	s.Push(MyInt(63))
	s.Pop()

	util.DD(s.Peek())

}
