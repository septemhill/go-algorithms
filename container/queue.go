package container

type Queue struct {
	l LinkedList
}

func (q *Queue) Enqueue(v int) {
	q.l.Prepend(v)
}

func (q *Queue) Dequeue() int {
	return q.l.RemoveTail()
}

func (q Queue) Trace() {
	q.l.Trace()
}
