package multipass

import (
	"errors"
	"testing"
)

func TestIsANew(t *testing.T) {
	ErrBar := NewError()
	err := error(ErrBar.New("foo"))
	if !ErrBar.IsA(err) {
		t.Error("err should be an ErrBar")
	}

	ErrFoo := NewError()
	if ErrFoo.IsA(err) {
		t.Error("err should not be an ErrFoo")
	}

	err = errors.New("baz")
	if ErrBar.IsA(err) {
		t.Error("err should not be an ErrBar")
	}
}

func TestEmbedCause(t *testing.T) {
	ErrFoo := NewError()
	ErrBar := errors.New("bar")
	err := ErrFoo.Embed(ErrBar, "foo")
	expected := "foo: bar"
	if err.Error() != expected {
		t.Error("Expected: %s, got: %s", expected, err.Error())
	}
	if err.Cause() != ErrBar {
		t.Error("err.Cause should be ErrBar")
	}

	ErrBaz := NewError()
	err = ErrBaz.Embed(err, "baz")
	expected = "baz: foo: bar"
	if err.Error() != expected {
		t.Error("Expected: %s, got: %s", expected, err.Error())
	}
	if err.Cause() != ErrBar {
		t.Error("err.Cause should be ErrBar")
	}
}

func TestError(t *testing.T) {
	ErrBar := NewError()
	expected := "foo"
	err := ErrBar.New(expected)
	if err.Error() != expected {
		t.Error("Expected %s, got: %s", expected, err.Error())
	}
}
