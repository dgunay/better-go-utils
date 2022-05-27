package logic

import (
	"fmt"
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

func (o Option[T]) OkOr(err error) Result[T] {
	if o.IsSome() {
		return Ok(o.value)
	}

	return Err[T](err)
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

func (o Option[T]) AndThen(fn func(val T) Option[T]) Option[T] {
	if o.IsSome() {
		return fn(o.value)
	}

	return None[T]()
}

func (o Option[T]) OrElse(fn func() Option[T]) Option[T] {
	if o.IsSome() {
		return o
	}

	return fn()
}
