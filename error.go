package multipass

import (
	"fmt"
	"sync/atomic"
)

var (
	nextErr = func(u uint64) *uint64 { return &u }(0)

	// enforce interfaces
	_ = error(Error{})
)

func id() uint64 { return atomic.AddUint64(nextErr, 1) }

// Error is an error type that supports more comparison, embedding and grouping.
type Error struct {
	id uint64

	Embedded error
	Text     string
}

// NewError returns a new Error that is a different kind from
// any other error.
func NewError() Error { return Error{id: id()} }

// IsA returns whether the specified error is the same kind as
// the current error.  To create errors of the same kind, use the
// New or Embed methods.
func (e Error) IsA(err error) bool {
	if otherError, ok := err.(Error); ok {
		return otherError.id == e.id
	}
	otherError, ok := err.(*Error)
	return ok && otherError.id == e.id
}

// New returns a new error of the same kind with the provided error message.
func (e Error) New(f string, a ...interface{}) Error {
	return Error{id: e.id, Text: fmt.Sprintf(f, a...)}
}

// Embed returns a new error of the same kind with the provided error message,
// that embeds the provided error.
func (e Error) Embed(err error, f string, a ...interface{}) Error {
	return Error{id: e.id, Embedded: err, Text: fmt.Sprintf(f, a...)}
}

// Error returns the consolidated error message for this error and all embedded errors.
func (e Error) Error() string {
	if e.Embedded == nil {
		return e.Text
	}
	return e.Text + ": " + e.Embedded.Error()
}

// Cause returns the root embedded error.
func (e Error) Cause() error {
	switch err := e.Embedded.(type) {
	case nil:
		return e
	case Error:
		return err.Cause()
	case *Error:
		return err.Cause()
	default:
		return err
	}
}
