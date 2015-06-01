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

type Error struct {
	id  uint64
	msg string
	e   error
}

func NewError() Error { return Error{id: id()} }
func (e Error) IsA(err error) bool {
	if otherError, ok := err.(Error); ok {
		return otherError.id == e.id
	}
	otherError, ok := err.(*Error)
	return ok && otherError.id == e.id
}
func (e Error) New(f string, a ...interface{}) Error {
	return Error{id: e.id, msg: fmt.Sprintf(f, a...)}
}
func (e Error) Embed(err error, f string, a ...interface{}) Error {
	return Error{id: e.id, e: err, msg: fmt.Sprintf(f, a...)}
}
func (e Error) Error() string {
	if e.e == nil {
		return e.msg
	}
	return e.msg + ": " + e.e.Error()
}
func (e Error) Cause() error {
	switch err := e.e.(type) {
	case nil:
		return e
	case Error:
		return err.Cause()
	case *Error:
		return err.Cause()
	default:
		return e.e
	}
}
