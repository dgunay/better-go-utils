package result

import (
	"fmt"
)

type Result[T any] struct {
	value T
	err   error
}

func Wrap[T any](args ...any) Result[T] {
	val := args[0].(T)
	err, ok := args[1].(error)
	if !ok {
		err = nil
	}

	return Result[T]{value: val, err: err}
}

func Ok[T any](val T) Result[T] {
	return Result[T]{value: val, err: nil}
}

func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func (r Result[T]) IsError() bool {
	return r.err != nil
}

func (r Result[T]) IsOk() bool {
	return !r.IsError()
}

func (r Result[T]) Unwrap() T {
	if r.IsError() {
		panic(fmt.Sprintf("unwrapped error result: %s", r.err))
	}

	return r.value
}

func (r Result[T]) Expect(msg string) T {
	if r.IsError() {
		panic(fmt.Sprintf("%s: %s", msg, r.err))
	}

	return r.value
}

func Map[T any, U any](r Result[T], fn func(val T) U) Result[U] {
	if r.IsOk() {
		return Ok(fn(r.value))
	}

	return Err[U](r.err)
}

// Mostly due to limitations in Go generics (methods can't have type params).
func (r Result[T]) DynMap(fn func(val T) any) Result[any] {
	if r.IsOk() {
		return Ok(fn(r.value))
	}

	return Result[any]{r.value, r.err}
}

func (r Result[T]) MapErr(fn func(err error) error) Result[T] {
	if r.IsError() {
		r.err = fn(r.err)
	}

	return r
}
