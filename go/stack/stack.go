package stack

type Stack[T any] struct {
	items []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Top() (T, bool) {
	var x T

	if len(s.items) > 0 {
		x = s.items[len(s.items)-1]
		return x, true
	}

	return x, false
}

func (s *Stack[T]) Pop() (T, bool) {
	var x T

	if len(s.items) > 0 {
		x, s.items = s.items[len(s.items)-1], s.items[:len(s.items)-1]

		return x, true
	}

	return x, false
}

func (s *Stack[T]) IsEmpty() bool {
	return (len(s.items) == 0)
}
