package collections_test

import (
	"testing"

	"github.com/dgunay/better-go-utils/collections"
	"github.com/stretchr/testify/assert"
)

func TestSliceReverse(t *testing.T) {
	t.Parallel()

	t.Run("reverses in place", func(t *testing.T) {
		t.Parallel()

		s := collections.NewSlice([]int{1, 2, 3})
		s.Reverse()
		assert.Equal(t, []int{3, 2, 1}, s.GetSlice())
	})
}

func TestSliceReversed(t *testing.T) {
	t.Parallel()

	t.Run("Returns a copy reversed, does not modify the original", func(t *testing.T) {
		t.Parallel()

		s := collections.NewSlice([]int{1, 2, 3})
		r := s.Reversed()
		assert.Equal(t, []int{3, 2, 1}, r.GetSlice())
		assert.Equal(t, []int{1, 2, 3}, s.GetSlice())
	})
}
