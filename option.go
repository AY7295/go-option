// Package option provides a generic implementation of the Option type, similar to Rust's Option type.
// It is used to encapsulate an optional value: a value that may either contain a value of type T or may be empty.
package option

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

// Nil is a pre-defined error to represent a nil value being incorrectly used to create an Option.
var (
	Nil = errors.New("nil value in Option")
)

// Option is a generic interface that defines methods for handling optional values.
type Option[T any] interface {
	Cause() error // Cause returns an error if the option is none due to an error.
	Ok() T        // Ok returns the value if present, otherwise a zero value.
}

// Some returns an Option containing the value t. t must be valid!
func Some[T any](t T) Option[T] {
	return &option[T]{
		val:   t,
		cause: nil,
	}
}

// None creates an Option that contains no value but an error.
// If no specific error is provided, it defaults to using Nil.
func None[T any](causes ...error) Option[T] {
	i := &option[T]{cause: errors.Join(causes...)}
	if i.cause == nil {
		i.cause = Nil
	}
	return i
}

// option is the internal struct implementing the Option interface for type T.
type option[T any] struct {
	val   T     // val holds the actual value.
	cause error // cause holds the error if the option is none.
}

// Cause returns the error cause for the option if it is none; otherwise, it returns nil.
func (opt *option[T]) Cause() error {
	return opt.cause
}

// Ok returns the stored value or a zero value if the option is none.
func (opt *option[T]) Ok() T {
	if opt.cause != nil {
		var zero T
		return zero
	}
	return opt.val
}

// IsSome returns true if the option holds a value (i.e., cause is nil).
func IsSome[T any](opt Option[T]) bool {
	return opt.Cause() == nil
}

// IsNone returns true if the option does not hold a value (i.e., cause is not nil).
func IsNone[T any](opt Option[T]) bool {
	return !IsSome(opt)
}

// Process applies a function to the value within the Option, returning a new Option with the result.
// If the original Option is none, the function is not executed and a new none Option with the original error is returned.
func Process[T, U any](opt Option[T], fn func(T) (U, error)) Option[U] {
	if opt.Cause() != nil {
		return None[U](fmt.Errorf("%v don't exec, pre-cause: %w",
			runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(),
			opt.Cause(),
		))
	}

	v, err := fn(opt.Ok())
	if err != nil {
		return None[U](err)
	}

	return Some(v)
}

// Map applies a function to the value within the Option, returning a new Option with the result.
// If the original Option is none, the function is not executed and a new none Option with the original error is returned.
func Map[T, U any](opt Option[T], fn func(T) U) Option[U] {
	if opt.Cause() != nil {
		return None[U](fmt.Errorf("%v don't exec, pre-cause: %w",
			runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(),
			opt.Cause(),
		))
	}

	return Some(fn(opt.Ok()))
}

// Flatten converts an Option[Option[T]] into a single Option[T].
// If the outer Option is none, it returns a none Option with the same error; otherwise, it returns the inner Option.
func Flatten[T any](opt Option[Option[T]]) Option[T] {
	if opt.Cause() != nil {
		return None[T](opt.Cause())
	}
	return opt.Ok()
}
