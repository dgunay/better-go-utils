package collections

// Any returns true if any element in the collection satisfies the predicate.
func Any[T any](elems []T, pred func(T) bool) bool {
	for _, elem := range elems {
		if pred(elem) {
			return true
		}
	}

	return false
}

// Map runs the given function on each element in the collection and returns a
// slice of the return values.
func Map[T any, U any](elems []T, mapfunc func(T) U) []U {
	result := make([]U, len(elems))

	for i, elem := range elems {
		result[i] = mapfunc(elem)
	}

	return result
}

func Filter[T any](elems []T, pred func(T) bool) []T {
	result := make([]T, 0)

	for _, elem := range elems {
		if pred(elem) {
			result = append(result, elem)
		}
	}

	return result
}

func ZeroOut[T any](elems []T) {
	for i := range elems {
		var zero T
		elems[i] = zero
	}
}

func Reverse[T any](elems []T) {
	for i, j := 0, len(elems)-1; i < j; i, j = i+1, j-1 {
		elems[i], elems[j] = elems[j], elems[i]
	}
}

func Reversed[T any](elems []T) []T {
	result := make([]T, len(elems))
	copy(result, elems)
	Reverse(result)
	return result
}
