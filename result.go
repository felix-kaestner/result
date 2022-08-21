package result

// Result discripes a discriminated union that encapsulates
// a successful outcome with a value of type T or a failure
// with an arbitrary error.
type Result[T any] interface {
	// Value returns the encapsulated value if
	// this instance represents success or the zero value
	// if it is failure.
	Value() T

	// Error returns the encapsulated error if this instance
	// represents failure or nil if it is success.
	Error() error

	// IsSuccess returns true if this instance represents
	// success or false if it is failure.
	IsSuccess() bool

	// IsFailure returns true if this instance represents
	// failure or false if it is success.
	IsFailure() bool

	// Fold executes the given functions when the instance
	// represents either a success or failure respectively.
	Fold(func(T), func(error))

	// OnSuccess executes the given function if the instance
	// represents a success.
	OnSuccess(func(T))

	// OnFailure executes the given function if the instance
	// represents a failure.
	OnFailure(func(error))

	// Recover executes the given function if the instance
	// represents a failure.
	// The encapsulated value will be set to the return value
	// of the function and the error to nil.
	Recover(func(error) T)
}

// result is an implementation of the Result interface.
type result[T any] struct {
	value T
	err   error
}

func (r *result[T]) Value() (t T) {
	return r.value
}

func (r *result[T]) Error() error {
	return r.err
}

func (r *result[T]) IsSuccess() bool {
	return r.err == nil
}

func (r *result[T]) IsFailure() bool {
	return r.err != nil
}

func (r *result[T]) Fold(onSuccess func(T), onFailure func(error)) {
	if r.IsSuccess() {
		onSuccess(r.value)
	} else {
		onFailure(r.err)
	}
}

func (r *result[T]) OnSuccess(fn func(T)) {
	if r.IsSuccess() {
		fn(r.value)
	}
}

func (r *result[T]) OnFailure(fn func(error)) {
	if r.IsFailure() {
		fn(r.err)
	}
}

func (r *result[T]) Recover(fn func(error) T) {
	if r.IsFailure() {
		r.value, r.err = fn(r.err), nil
	}
}

// New creates a new Result instance representing either a success or failure
// depending on the arguments.
func New[T any](v T, err error) Result[T] {
	return &result[T]{
		value: v,
		err:   err,
	}
}

// Success creates a new Result instance representing a success.
func Success[T any](v T) Result[T] {
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
