package promise

import (
	"context"

	"github.com/dgunay/better-go-utils/logic"
)

type Promise[T any] struct {
	channel chan T
	errch   chan error
	ready   bool
	cancel  context.CancelFunc
}

// The passed context's Done channel is used to cancel the promise.
type PromiseFunc[T any] func(context.Context) (T, error)

func New[T any](fn PromiseFunc[T]) *Promise[T] {
	p := &Promise[T]{channel: make(chan T, 1), errch: make(chan error, 1)}

	ctx := context.Background()
	ctx, p.cancel = context.WithCancel(ctx)
	go func() {
		val, err := fn(ctx)
		if err != nil {
			p.errch <- err
		} else {
			p.channel <- val
		}
		p.ready = true
	}()

	return p
}

func (p Promise[T]) Ready() logic.Bool {
	return logic.Bool(p.ready)
}

func (p Promise[T]) Await() (T, error) {
	for {
		select {
		case val := <-p.channel:
			return val, nil
		case err := <-p.errch:
			var none T
			return none, err
		}
	}
}

// Sends a cancellation/Done() to the context passed to the PromiseFunc. It is
// up to the callback to respond to the cancellation.
func (p Promise[T]) Cancel() {
	p.cancel()
}
