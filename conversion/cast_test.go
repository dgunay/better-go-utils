package conversion_test

import (
	"testing"

	. "github.com/dgunay/better-go-utils/conversion"
	"github.com/dgunay/better-go-utils/result"
	"github.com/stretchr/testify/assert"
)

type someInterface interface {
	Do() string
}

type A struct{}

func (a A) Do() string { return "A" }

type B struct{}

func (b B) Do() string { return "B" }

func TestCast(t *testing.T) {
	t.Parallel()

	t.Run("Can cast to/from the same type", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, result.Ok(1), Cast[int](1))
	})

	t.Run("Can cast interfaces into their underlying implementor", func(t *testing.T) {
		t.Parallel()

		interfA := someInterface(&A{})
		interfB := someInterface(&B{})

		concreteA := Cast[*A](interfA).Unwrap()
		concreteB := Cast[*B](interfB).Unwrap()

		assert.Equal(t, "A", concreteA.Do())
		assert.Equal(t, "B", concreteB.Do())
	})

	t.Run("cannot cast incompatible types", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, result.Err[int](ErrInvalidCast("", 1)), Cast[int](""))
	})
}
