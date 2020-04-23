package queue

type Node struct {
	value interface{}
	next  *Node
}

type Queue struct {
	root  *Node
	tail *Node
	_size uint
}

func (q *Queue) Top() *Node {
	if q._size > 0 {
		return q.root
	}
	return nil
}

func (q *Queue) Pop() {
	if q._size > 0 {
		q.root = q.root.next
		q._size--
	}
}

func (q *Queue) Push(e interface{}){
	q.tail.next = &Node{e, nil}
	q._size++
}

func (q *Queue) Size() uint {
	return q._size
}