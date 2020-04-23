package queue

type Node struct {
	value interface{}
	next  *Node
}

type Queue struct {
	root  *Node
	tail  *Node
	_size uint
}

func (q *Queue) Top() interface{} {
	if q._size > 0 {
		return q.root.value
	}
	return nil
}

func (q *Queue) Pop() {
	if q._size > 0 {
		q.root = q.root.next
		q._size--
	}
}

func (q *Queue) Push(e interface{}) {
	t :=  &Node{e, nil}
	if q.root == nil {
		q.root = t
		q.tail = t
	}else{
		q.tail.next = t
		q.tail = q.tail.next
	}
	q._size++
}

func (q *Queue) Size() uint {
	return q._size
}
