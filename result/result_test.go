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
		resInt := result.Ok(1)

		resString := result.Map(resInt, func(val int) string { return fmt.Sprintf("%d", val) })
		require.Equal(t, "1", resString.Unwrap())

		backToInt := result.Map(resString, func(val string) any {
			return result.Wrap[int](strconv.Atoi(val)).Unwrap()
		})

		require.Equal(t, 1, backToInt.Unwrap())

	})
}
