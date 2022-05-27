package collections

type Stack[T any] struct {
	Slice[T]
}

// Alias for Append
func (s *Slice[T]) Push(elem T) {
	s.Append(elem)
}

func (s *Stack[T]) Pop() T {
	elem := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return elem
}
