package collections

import "github.com/dgunay/better-go-utils/option"

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

func (s Slice[T]) Any(pred func(T) bool) bool {
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

func (s Slice[T]) Last() option.Option[T] {
	if s.Empty() {
		return option.None[T]()
	}

	return s.At(s.Len() - 1)
}

func (s Slice[T]) First() option.Option[T] {
	if s.Empty() {
		return option.None[T]()
	}

	return option.Some(s.data[0])
}

func (s Slice[T]) Empty() bool {
	return s.Len() == 0
}

func (s Slice[T]) At(i int) option.Option[T] {
	if i < 0 { // negatives wrap around
		i = s.Len() + i
	}

	if i < 0 || i >= s.Len() {
		return option.None[T]()
	}

	return option.Some(s.data[i])
}

// TODO: test that this actually deep copies
// func (s Slice[T]) Copy() Slice[T] {
// 	return NewSlice(s.data)
// }
