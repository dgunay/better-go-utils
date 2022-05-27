package conversion

import (
	"reflect"

	"github.com/dgunay/better-go-utils/result"
)

type InvalidCastError struct {
	FromType string
	ToType   string
}

func ErrInvalidCast[T any, U any](from T, to U) error {
	return InvalidCastError{
		FromType: reflect.TypeOf(from).String(),
		ToType:   reflect.TypeOf(to).String(),
	}
}

func (e InvalidCastError) Error() string {
	return "Invalid cast from " + e.FromType + " to " + e.ToType
}

func Cast[Target any, From any](from From) result.Result[Target] {
	casted, ok := any(from).(Target)
	if !ok {
		var zeroTarget Target
		return result.Err[Target](ErrInvalidCast(from, zeroTarget))
	}

	return result.Ok(casted)
}
