package collections_test

import (
	"testing"

	"github.com/dgunay/better-go-utils/collections"
	"github.com/dgunay/better-go-utils/option"
	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	t.Parallel()
	t.Run("appends element to the underlying slice", func(t *testing.T) {
		s := collections.NewStack([]int{1, 2, 3})

		s.Push(4)
		assert.Equal(t, 4, s.Len())
	})

	t.Run("can push multiple elements", func(t *testing.T) {
		s := collections.NewStack([]int{1, 2, 3})

		s.Push(4, 5)
		assert.Equal(t, 5, s.Len())
		assert.Equal(t, 5, s.Pop().Unwrap())
		assert.Equal(t, 4, s.Pop().Unwrap())
	})
}

func TestPop(t *testing.T) {
	t.Parallel()
	t.Run("Return the last value as Some", func(t *testing.T) {
		t.Parallel()

		s := collections.NewStack([]int{1, 2, 3})
		assert.Equal(t, 3, s.Len())
		last := s.Pop()
		assert.Equal(t, 2, s.Len())
		assert.Equal(t, 3, last.Unwrap())
	})

	t.Run("Returns None if empty", func(t *testing.T) {
		t.Parallel()
		s := collections.NewStack([]int{})
		assert.Equal(t, 0, s.Len())
		last := s.Pop()
		assert.Equal(t, 0, s.Len())
		assert.Equal(t, option.None[int](), last)
	})
}
