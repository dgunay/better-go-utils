package result

import (
	"fmt"
)

type Result[T any, MapTo any] struct {
	value T
	err   error
}

func Wrap[T any, M any](args ...any) Result[T, M] {
	val := args[0].(T)
	err, ok := args[1].(error)
	if !ok {
		err = nil
	}

	return Result[T, M]{value: val, err: err}
}

func Ok[T any, M any](val T) Result[T, M] {
	return Result[T, M]{value: val, err: nil}
}

func Err[T any, M any](err error) Result[T, M] {
	return Result[T, M]{err: err}
}

func (r Result[T, M]) IsError() bool {
	return r.err != nil
}

func (r Result[T, M]) IsOk() bool {
	return !r.IsError()
}

func (r Result[T, M]) Unwrap() T {
	if r.IsError() {
		panic(fmt.Sprintf("unwrapped error result: %s", r.err))
	}

	return r.value
}

func (r Result[T, M]) Expect(msg string) T {
	if r.IsError() {
		panic(fmt.Sprintf("%s: %s", msg, r.err))
	}

	return r.value
}

func (r Result[T, M]) Map(fn func(val T) M) Result[M, any] {
	if r.IsOk() {
		return Ok[M, any](fn(r.value))
	}

	return Err[M, any](r.err)
}

// Mostly due to limitations in Go generics (methods can't have type params).
func (r Result[T, M]) DynMap(fn func(val T) any) Result[any, any] {
	if r.IsOk() {
		return Ok[any, any](fn(r.value))
	}

	return Result[any, any]{r.value, r.err}
}

func (r Result[T, M]) MapErr(fn func(err error) error) Result[T, M] {
	if r.IsError() {
		r.err = fn(r.err)
	}

	return r
}

func (r Result[T, M]) AndThen(fn func(val T) M) Result[M, any] {
	if r.IsOk() {
		return r.Map(fn)
	}

	return Err[M, any](r.err)
}
