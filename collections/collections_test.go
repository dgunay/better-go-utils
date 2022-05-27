package collections_test

import (
	"testing"

	"github.com/dgunay/better-go-utils/collections"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	t.Parallel()

	t.Run("Reverses a slice in-place", func(t *testing.T) {
		t.Parallel()

		s := []int{1, 2, 3}
		collections.Reverse(s)
		assert.Equal(t, []int{3, 2, 1}, s)
	})
}

func TestReversed(t *testing.T) {
	t.Parallel()

	t.Run("Returns a reversed slice, does not modify the original", func(t *testing.T) {
		t.Parallel()

		s := []int{1, 2, 3}
		r := collections.Reversed(s)
		assert.Equal(t, []int{3, 2, 1}, r)
		assert.Equal(t, []int{1, 2, 3}, s)
	})
}
