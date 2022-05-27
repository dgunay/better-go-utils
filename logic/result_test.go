package logic_test

import (
	"fmt"
	"strconv"
	"testing"

	. "github.com/dgunay/better-go-utils/logic"
	"github.com/stretchr/testify/require"
)

func TestMapping(t *testing.T) {
	t.Run("Mapping works", func(t *testing.T) {
		resInt := Ok(1)

		resString := Map(resInt, func(val int) string { return fmt.Sprintf("%d", val) })
		require.Equal(t, "1", resString.Unwrap())

		backToInt := Map(resString, func(val string) any {
			return Wrap[int](strconv.Atoi(val)).Unwrap()
		})

		require.Equal(t, 1, backToInt.Unwrap())
	})
}
