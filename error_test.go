package multipass

import (
	"errors"
	"fmt"
	"testing"
)

func ExampleError_IsA() {
	ErrFoo := NewError()

	err := error(ErrFoo.New("foo"))
	fmt.Println(ErrFoo.IsA(err))

	err = errors.New("bar")
	fmt.Println(ErrFoo.IsA(err))

	// Output: true
	// false

}

func ExampleError_Cause() {
	ErrFoo := NewError()
	ErrBar := NewError()
	ErrBaz := errors.New("baz")

	err := ErrBar.Embed(ErrBaz, "bar")
	err = ErrFoo.Embed(err, "foo")
	fmt.Println(err.Cause() == ErrBaz)

	// Output: true
}

func ExampleError_Error() {
	ErrFoo := NewError()
	ErrBar := NewError()
	ErrBaz := errors.New("baz")

	err := ErrBar.Embed(ErrBaz, "bar")
	err = ErrFoo.Embed(err, "foo")
	fmt.Println(err.Error())

	// Output foo: bar: baz
}

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
