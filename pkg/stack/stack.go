package stack

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Stack[T any] struct {
	head *Node[T]
	size uint
}

func New[T any](values ...T) *Stack[T] {
	s := &Stack[T]{}

	for i := range values {
		s.Push(values[len(values)-i-1])
	}

	return s
}

func (s *Stack[T]) Push(value T) {
	n := &Node[T]{value, s.head}
	s.head = n
	s.size++
}

func (s *Stack[T]) Pop() T {
	if s.head == nil {
		panic("Stack is empty.")
	}

	n := s.head
	s.head = n.next
	s.size--

	return n.value
}

func (s *Stack[T]) GetSize() uint {
	return s.size
}
