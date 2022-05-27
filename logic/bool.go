package logic

type Bool bool

// Can't implement this yet because of generics limitations
func (b Bool) ThenSomeAny(thing any) Option[any] {
	if b {
		return Some(thing)
	}

	return None[any]()
}

func (b Bool) Then(f func() Bool) Bool {
	if b {
		return f()
	}

	return b
}

func (b Bool) Else(f func() Bool) Bool {
	if !b {
		return f()
	}

	return b
}
