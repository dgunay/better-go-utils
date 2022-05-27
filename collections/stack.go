package collections

import "github.com/dgunay/better-go-utils/option"

type Stack[T any] struct {
	slice Slice[T]
}

func NewStack[T any](data []T) Stack[T] {
	return Stack[T]{NewSlice(data)}
}

func (s Stack[T]) Len() int {
	return s.slice.Len()
}

// Alias for Append
func (s *Stack[T]) Push(elem T) {
	s.slice.Append(elem)
}

func (s *Stack[T]) Pop() option.Option[T] {
	return s.slice.Last().AndThen(func(val T) option.Option[T] {
		s.slice.data = s.slice.data[:len(s.slice.data)-1]
		return option.Some(val)
	})
}
