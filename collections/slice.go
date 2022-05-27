package collections

import "github.com/dgunay/better-go-utils/logic"

// Slice wraps a slice and exposes convenient methods for working with it.
type Slice[T any] struct {
	data []T
}

func NewSlice[T any](data []T) Slice[T] {
	return Slice[T]{data}
}

func (s Slice[T]) GetSlice() []T {
	return s.data
}

func (s Slice[T]) Len() int {
	return len(s.data)
}

func (s Slice[T]) Any(pred func(T) bool) logic.Bool {
	return Any(s.data, pred)
}

// Go generics don't have method type parameters, so this is the best we can
// do for now.
func (s Slice[T]) DynMap(mapfunc func(T) any) Slice[any] {
	return NewSlice(Map(s.data, mapfunc))
}

func (s Slice[T]) Filter(pred func(T) bool) Slice[T] {
	return NewSlice(Filter(s.data, pred))
}

func (s Slice[T]) ForEach(fn func(T)) {
	for _, elem := range s.data {
		fn(elem)
	}
}

// Allows you to directly access the address of the underlying slice elements.
func (s *Slice[T]) ForEachByRef(fn func(*T)) {
	for i := range s.data {
		fn(&s.data[i])
	}
}

// Replaces all of the elements of the slice with the zero value of the type.
func (s *Slice[T]) Zeroed() {
	ZeroOut(s.data)
}

func (s *Slice[T]) Append(elem ...T) {
	s.data = append(s.data, elem...)
}

func (s Slice[T]) Last() logic.Option[T] {
	if s.Empty() {
		return logic.None[T]()
	}

	return s.At(s.Len() - 1)
}

func (s Slice[T]) First() logic.Option[T] {
	if s.Empty() {
		return logic.None[T]()
	}

	return logic.Some(s.data[0])
}

func (s Slice[T]) Empty() logic.Bool {
	return s.Len() == 0
}

func (s Slice[T]) At(i int) logic.Option[T] {
	if i < 0 { // negatives wrap around
		i = s.Len() + i
	}

	if i < 0 || i >= s.Len() {
		return logic.None[T]()
	}

	return logic.Some(s.data[i])
}

func (s Slice[T]) Reversed() Slice[T] {
	return NewSlice(Reversed(s.data))
}

func (s *Slice[T]) Reverse() {
	Reverse(s.data)
}
