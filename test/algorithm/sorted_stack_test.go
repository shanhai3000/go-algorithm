package algorithm

import (
	"algo/algorithm"
	"algo/algorithm/util"
	"testing"
)

func TestSortedStack(t *testing.T) {
	s := new(algorithm.SortedStack)
	s.Pop()
	s.Push(40)
	s.Pop()
	s.Push(19)
	s.Push(44)
	s.Pop()
	s.Pop()
	s.Push(42)
	s.Push(8)
	s.Push(29)
	s.Push(25)
	s.Pop()
	s.Pop()
	s.Push(52)
	s.Push(63)
	s.Pop()
	s.Push(47)
	s.Pop()
	s.Push(45)
	s.Push(52)
	s.Pop()
	s.Pop()
	s.Push(17)
	s.Pop()
	s.Push(6)
	s.Push(30)
	s.Push(51)
	s.Push(46)
	s.Push(2)
	s.Push(56)
	s.Push(39)
	s.Push(38)

	util.DD(s.Peek())

}
