package result_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/dgunay/better-go-utils/result"
	"github.com/stretchr/testify/require"
)

func TestMapping(t *testing.T) {
	t.Run("Mapping works", func(t *testing.T) {
		resInt := result.Ok[int, string](1)

		resString := resInt.Map(func(val int) string { return fmt.Sprintf("%d", val) })
		require.Equal(t, "1", resString.Unwrap())

		backToInt := resString.Map(func(val string) any {
			return result.Wrap[int, any](strconv.Atoi(val)).Unwrap()
		})

		require.Equal(t, 1, backToInt.Unwrap())

	})
}
