package result

// Result discripes a discriminated union that encapsulates
// a successful outcome with a value of type T or a failure
// with an arbitrary error.
type Result[T any] interface {
	// Value returns the encapsulated value if this instance
	// represents success or nil if it is failure.
	Value() *T

	// Error returns the encapsulated error if this instance
	// represents failure or nil if it is success.
	Error() error

	// IsSuccess returns true if this instance represents
	// success or false if it is failure.
	IsSuccess() bool

	// IsFailure returns true if this instance represents
	// failure or false if it is success.
	IsFailure() bool
}

// result is an implementation of the Result interface.
type result[T any] struct {
	value *T
	err   error
}

func (r *result[T]) Value() *T {
	return r.value
}

func (r *result[T]) Error() error {
	return r.err
}

func (r *result[T]) IsSuccess() bool {
	return r.value != nil
}

func (r *result[T]) IsFailure() bool {
	return r.err != nil
}

// Success creates a new Result instance representing a success.
func Success[T any](v *T) Result[T] {
	return &result[T]{
		value: v,
	}
}

// Failure creates a new Result instance representing a failure.
func Failure[T any](err error) Result[T] {
	return &result[T]{
		err: err,
	}
}
