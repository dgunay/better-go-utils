package promise_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/dgunay/better-go-utils/promise"
	"github.com/stretchr/testify/require"
)

func TestPromise(t *testing.T) {
	success := func(context.Context) (string, error) {
		return "success", nil
	}

	failure := func(context.Context) (string, error) {
		return "", fmt.Errorf("failure")
	}

	t.Run("happy path", func(t *testing.T) {
		p := promise.New(success)
		require.Eventually(t, func() bool { return bool(p.Ready()) }, time.Second, time.Millisecond*50)
		val, err := p.Await()
		require.NoError(t, err)
		require.Equal(t, "success", val)
	})

	t.Run("error path", func(t *testing.T) {
		p := promise.New(failure)
		require.Eventually(t, func() bool { return bool(p.Ready()) }, time.Second, time.Millisecond*50)
		val, err := p.Await()
		require.Error(t, err)
		require.Equal(t, "", val)
	})

	t.Run("cancelling the promise", func(t *testing.T) {
		failWhenCanceled := func(ctx context.Context) (string, error) {
			for {
				select {
				case <-ctx.Done():
					return "", fmt.Errorf("cancelled")
				case <-time.After(time.Second):
					return "success", nil
				}
			}
		}

		// Cancel the promise
		p := promise.New(failWhenCanceled)
		p.Cancel()
		require.Eventually(t, func() bool { return bool(p.Ready()) }, time.Second*2, time.Millisecond*50)
		val, err := p.Await()
		require.ErrorContains(t, err, "cancelled")
		require.Equal(t, "", val)

		// Don't cancel
		p = promise.New(failWhenCanceled)
		require.Eventually(t, func() bool { return bool(p.Ready()) }, time.Second*2, time.Millisecond*50)
		val, err = p.Await()
		require.NoError(t, err)
		require.Equal(t, "success", val)
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
