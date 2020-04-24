package collection

import (
	"algo/collection/queue"
	"testing"
)

func TestQueue(t *testing.T){
	q := new(queue.Queue)
	q.Push(1)
	if q.Size() == 1{
		t.Log("Size() ok", q.Size())
	}else{
		t.Error("Size() error", q.Size())
	}

	e := q.Top()
	if e == 1 {
		t.Log("Top() ok", e)
	}else{
		t.Error("Top() error")
	}
	q.Pop()
	if q.Size() == 0 {
		t.Log("Pop() ok")
	}else{
		t.Error("Pop() error")
	}
}
