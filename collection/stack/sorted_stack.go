package stack

import (
	"algo/collection/node"
	"algo/collection/sort"
)

type (
	SortedStack struct {
		Root *node.Node
		Mix  *node.Node
	}
)

func (s *SortedStack) Push(val sort.Comparable) {
	t := &node.Node{Value: val}
	if s.Root == nil {
		s.Root = t
	} else {
		if s.Root.Value.(sort.Comparable).CompareTo(val) >= 0 {
			t.Next = s.Root
			s.Root = t
		} else {
			if s.Mix == nil || s.Mix.Value.(sort.Comparable).CompareTo(val) >= 0 {
				t.Next = s.Mix
				s.Mix = t
			} else {
				//比两个栈的栈顶都大，需要放入主栈
				_r, _m := s.Root, s.Mix
				for _r != nil && _r.Value.(sort.Comparable).CompareTo(val) < 0 {
					newR := _r.Next
					_r.Next = _m
					_m = _r
					_r = newR
				}
				if _r == nil {
					_r = t
				} else {
					t.Next = _r
					_r = t
				}
				var __m *node.Node
				var newM *node.Node
				for ; _m != nil && _r.Value.(sort.Comparable).CompareTo(_m.Value.(sort.Comparable)) >= 0; {
					__m = _m
					for ; __m.Next != nil && __m.Next.Value.(sort.Comparable).CompareTo(_r.Value.(sort.Comparable)) <= 0 && __m.Next.Value.(sort.Comparable).CompareTo(__m.Value.(sort.Comparable)) >= 0; __m = __m.Next {
					}
					newM = __m.Next
					__m.Next = _r
					_r = _m
					_m = newM
				}
				s.Root = _r
				s.Mix = _m
			}
		}
	}
}

func (s *SortedStack) Pop() {
	if s.Root != nil {
		s.Root = s.Root.Next
	}
	_r, _m := s.Root, s.Mix
	if _m != nil {
		__m := _m
		if _r == nil {
			for ; __m.Next != nil && __m.Next.Value.(sort.Comparable).CompareTo(_m.Value.(sort.Comparable)) >= 0; __m = __m.Next {
			}
			newM := __m.Next
			__m.Next = nil
			_r = _m
			_m = newM
		} else {
			if _r.Value.(sort.Comparable).CompareTo(_m.Value.(sort.Comparable)) >= 0 {
				__m := _m
				for ; __m.Next != nil && __m.Next.Value.(sort.Comparable).CompareTo(_r.Value.(sort.Comparable)) <= 0 && __m.Next.Value.(sort.Comparable).CompareTo(_m.Value.(sort.Comparable)) >= 0; __m = __m.Next {
				}
				tmpM := __m.Next
				__m.Next = _r
				_r = _m
				_m = tmpM
			}
		}
	}
	s.Root = _r
	s.Mix = _m
}

func (s *SortedStack) Peek() sort.Comparable {
	if s.Root != nil {
		return s.Root.Value.(sort.Comparable)
	}
	return nil
}

func (s *SortedStack) Empty() bool {
	return s.Root == nil
}
