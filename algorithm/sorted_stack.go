package algorithm

import "algo/collection/node"

type (
	SortedStack struct {
		Root *node.Node
		Mix  *node.Node
	}
)

func (s *SortedStack) Push(val int) {
	t := &node.Node{Val: val}
	if s.Root == nil {
		s.Root = t
	} else {
		if s.Root.Val >= val {
			t.Next = s.Root
			s.Root = t
		} else {
			if s.Mix == nil || val <= s.Mix.Val {
				t.Next = s.Mix
				s.Mix = t
			} else {
				//比两个栈的栈顶都大，需要放入主栈
				_r, _m := s.Root, s.Mix
				for _r != nil && _r.Val < val {
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
				for ; _m != nil && _r.Val >= _m.Val; {
					__m = _m
					for ; __m.Next != nil && __m.Next.Val <= _r.Val && __m.Next.Val >= __m.Val; __m = __m.Next {
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
			for ; __m.Next != nil && __m.Next.Val >= _m.Val; __m = __m.Next {
			}
			newM := __m.Next
			__m.Next = nil
			_r = _m
			_m = newM
		} else {
			if _r.Val >= _m.Val {
				__m := _m
				for ; __m.Next != nil && __m.Next.Val <= _r.Val && __m.Next.Val >= _m.Val; __m = __m.Next {
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

func (s *SortedStack) Peek() int {
	if s.Root != nil {
		return s.Root.Val
	}
	return -1
}

func (s *SortedStack) IsEmpty() bool {
	return s.Root == nil
}
