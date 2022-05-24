package promise_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgunay/better-go-utils/promise"
	"github.com/stretchr/testify/require"
)

func TestPromise(t *testing.T) {
	success := func() (string, error) {
		return "success", nil
	}

	failure := func() (string, error) {
		return "", fmt.Errorf("failure")
	}

	t.Run("happy path", func(t *testing.T) {
		p := promise.New(success)
		require.Eventually(t, func() bool { return p.Ready() }, time.Second, time.Millisecond*50)
		val, err := p.Await()
		require.NoError(t, err)
		require.Equal(t, "success", val)
	})

	t.Run("error path", func(t *testing.T) {
		p := promise.New(failure)
		require.Eventually(t, func() bool { return p.Ready() }, time.Second, time.Millisecond*50)
		val, err := p.Await()
		require.Error(t, err)
		require.Equal(t, "", val)
	})

	// TODO: implement chaining thens
	// t.Run("chain with .Then", func(t *testing.T) {
	// 	var called *bool = new(bool)
	// 	p := promise.New(success).Then(func(val string) error {
	// 		require.Equal(t, "success", val)
	// 		*called = true
	// 		return nil
	// 	})
	// 	require.Eventually(t, func() bool { return p.Ready() }, time.Second, time.Millisecond*50)
	// 	val, err := p.Await()
	// 	require.NoError(t, err)
	// 	require.Equal(t, "success", val)
	// 	require.True(t, *called)
	// })
}
