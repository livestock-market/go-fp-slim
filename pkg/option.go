package gofp

import (
	"encoding/json"
)

// Option represents a value that may or may not be present.
type Option[T any] struct {
	value    T
	isEmpty  bool
	hasValue bool
}

// Some returns a new Option with the given value.
func Some[T any](value T) Option[T] {
	return Option[T]{
		value:    value,
		isEmpty:  false,
		hasValue: true,
	}
}

// None returns an empty Option.
func None[T any]() Option[T] {
	return Option[T]{
		isEmpty:  true,
		hasValue: false,
	}
}

// FromNilable returns an Option with the given value if it is not nil,
// otherwise returns an empty Option.
func FromNilable[T any](value *T) Option[T] {
	if IsNil(value) {
		return None[T]()
	}
	return Some(*value)
}

// IsEmpty returns true if the Option is empty, false otherwise.
func (o Option[T]) IsEmpty() bool {
	return o.isEmpty
}

// HasValue returns true if the Option has a value, false otherwise.
func (o Option[T]) HasValue() bool {
	return o.hasValue
}

// Get returns the value of the Option.
// **Panics if the Option is empty.**
func (o Option[T]) Get() T {
	if o.isEmpty {
		panic("Get(): No value present, use IsEmpty() to check")
	}
	return o.value
}

// Implement the json.Marshaler interface
func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.isEmpty {
		// For empty Option, return null unless JSON v2 is released
		// with omitting empty fields
		return []byte("null"), nil
	}
	// For non-empty Option, serialize the value
	return json.Marshal(o.value)
}
