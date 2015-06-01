package multipass

import (
	"errors"
	"testing"
)

func TestContains(t *testing.T) {
	ErrFoo := NewError()
	ErrBar := NewError()
	ErrBaz := errors.New("baz")
	ErrQux := errors.New("qux")
	ErrQuux := NewError()

	ErrClass := NewErrorClass(ErrFoo, ErrBar, ErrBaz)

	if !ErrClass.Contains(ErrBar) {
		t.Error("ErrBar is a member of ErrClass")
	}
	if !ErrClass.Contains(ErrBaz) {
		t.Error("ErrBaz is a member of ErrClass")
	}
	if ErrClass.Contains(ErrQux) {
		t.Error("ErrQux is not a member of ErrClass")
	}
	if ErrClass.Contains(ErrQuux) {
		t.Error("ErrQuux is not a member of ErrClass")
	}
}
