package container

type Stack struct {
	l LinkedList
}

func (s *Stack) Pop() int {
	return s.l.RemoveHead()
}

func (s *Stack) Push(v int) {
	s.l.Prepend(v)
}

func (s Stack) Trace() {
	s.l.Trace()
}
