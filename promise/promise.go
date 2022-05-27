package promise

import "github.com/dgunay/better-go-utils/logic"

type Promise[T any] struct {
	channel chan T
	errch   chan error
	ready   bool
}

func New[T any](fn func() (T, error)) *Promise[T] {
	p := &Promise[T]{channel: make(chan T, 1), errch: make(chan error, 1)}

	go func() {
		val, err := fn()
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
