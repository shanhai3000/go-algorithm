package queue

type (
	node struct {
		value interface{}
		next  *node
	}

	Queue struct {
		root  *node
		tail  *node
		_size uint
	}
)

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
		if q._size == 0 {
			q.tail = nil
		}
	}
}

func (q *Queue) Push(e interface{}) {
	t := &node{e, nil}
	if q.root == nil {
		q.root = t
		q.tail = t
	} else {
		q.tail.next = t
		q.tail = q.tail.next
	}
	q._size++
}

func (q *Queue) Size() uint {
	return q._size
}
