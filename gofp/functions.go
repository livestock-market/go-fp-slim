package gofp

// Returns true if the given pointer points to a nil, false otherwise.
func IsNil[A any](a *A) bool {
	return a == nil
}

// Returns true if the given pointer points to a non-nil value, false otherwise.
func IsNonNil[A any](a *A) bool {
	return a != nil
}

// Applies the given predicate function to the given slice and returns a new slice
// containing only the elements that satisfy the predicate.
func Filter[T any](source []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range source {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

const (
	zero = 0
)

// Applies the given function to the given slice and returns a new slice containing the results.
func Map[T, U any](f func(T) U, arr []T) []U {
	if len(arr) == zero {
		return []U{}
	}

	result := make([]U, len(arr))
	for i, item := range arr {
		result[i] = f(item)
	}
	return result
}
