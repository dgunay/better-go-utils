package option

import (
	"fmt"

	"github.com/dgunay/better-go-utils/result"
)

type Option[T any] struct {
	value  T
	isSome bool
}

func FromPtr[T any](pointer *T) Option[T] {
	if pointer == nil {
		return None[T]()
	}

	return Some(*pointer)
}

func Some[T any](val T) Option[T] {
	return Option[T]{value: val, isSome: true}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func (o Option[T]) IsSome() bool {
	return o.isSome
}

func (o Option[T]) Unwrap() T {
	if !o.IsSome() {
		panic("unwrapped None")
	}

	return o.value
}

func (o Option[T]) Expect(msg string) T {
	if !o.IsSome() {
		panic(fmt.Sprintf("%s: unwrapped None", msg))
	}

	return o.value
}

func (o Option[T]) OkOr(err error) result.Result[T] {
	if o.IsSome() {
		return result.Ok(o.value)
	}

	return result.Err[T](err)
}

// Mostly due to limitations in Go generics (methods can't have type params).
func (o Option[T]) DynMap(fn func(val T) any) Option[any] {
	if o.IsSome() {
		return Some(fn(o.value))
	}

	return None[any]()
}

func (o *Option[T]) Take() Option[T] {
	out := *o

	if o.IsSome() {
		out = Some(o.value)
		*o = None[T]()
	}

	return out
}
